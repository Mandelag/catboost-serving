#include <chrono>
#include <iostream>
#include <seastar/core/app-template.hh>
#include <seastar/core/print.hh>
#include <seastar/core/prometheus.hh>
#include <seastar/core/reactor.hh>
#include <seastar/core/seastar.hh>
#include <seastar/core/sleep.hh>
#include <seastar/core/thread.hh>
#include <seastar/http/api_docs.hh>
#include <seastar/http/file_handler.hh>
#include <seastar/http/function_handlers.hh>
#include <seastar/http/handlers.hh>
#include <seastar/http/httpd.hh>
#include <seastar/net/api.hh>
#include <seastar/core/iostream.hh>
#include <seastar/core/iostream-impl.hh>


void routes(seastar::httpd::routes &r) {
  seastar::httpd::function_handler *h1 = new seastar::httpd::function_handler(
      [](seastar::const_req req) { return "hello\n"; });
  r.add(seastar::operation_type::GET, seastar::httpd::url("/"), h1);
}

int main(int argc, char **argv) {
  seastar::app_template app;
  app.run(argc, argv, [] {
    seastar::socket_address address =
          seastar::socket_address(seastar::ipv4_addr("0.0.0.0:9000"));
    seastar::listen_options listen_opt = {reuse_address : true};
    auto server = new seastar::httpd::http_server_control();
    std::cout << "Starting server..\n"; 
    
    auto th2 = seastar::async([server, &address, &listen_opt] {
      server->start("Alamancay server").wait();
      server->set_routes(routes).wait();
      server->listen(address, listen_opt)
          .then([] { std::cout << "Alamantap?\n"; })
          .wait();
      std::cout << "HTTP Server ready on 0.0.0.0:9000 ...\n";
    });
    
    
    
    return seastar::async([server] {
      // thread for signaling..

      using namespace std::chrono_literals;
      // TODO: obviously
      seastar::sleep_abortable(std::chrono::hours(9999999999999999)).wait();
      std::cout << "Exiting....\n";
      server->stop().wait();
      std::cout << "Exited!\n";
    });
  });
}
