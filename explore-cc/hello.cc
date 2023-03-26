#include <seastar/core/print.hh>
#include <seastar/core/prometheus.hh>
#include <seastar/core/reactor.hh>
#include <seastar/core/seastar.hh>
#include <seastar/core/thread.hh>
#include <seastar/http/api_docs.hh>
#include <seastar/http/file_handler.hh>
#include <seastar/http/function_handlers.hh>
#include <seastar/http/handlers.hh>
#include <seastar/http/httpd.hh>

void routes(seastar::httpd::routes &r) {
  seastar::httpd::function_handler *h1 = new seastar::httpd::function_handler(
      [](seastar::const_req req) { return "hello"; });
  r.add(seastar::operation_type::GET, seastar::httpd::url("/"), h1);
}

int main(int argc, char **argv) {
  seastar::httpd::http_server_control server;
  seastar::socket_address address =
      seastar::socket_address(seastar::ipv4_addr("0.0.0.0:9000"));

  seastar::app_template app;
  return app.run_deprecated(argc, argv, [&] {
    return seastar::async([&] {
      server.start("my-first-seastar-http-server")
          .then([&server] { return server.set_routes(routes); })
          .then([&server, &address] {
            return server.listen(address).then([] {});
          });
    });
  });
}
