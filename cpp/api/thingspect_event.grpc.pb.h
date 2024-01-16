// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: api/thingspect_event.proto
#ifndef GRPC_api_2fthingspect_5fevent_2eproto__INCLUDED
#define GRPC_api_2fthingspect_5fevent_2eproto__INCLUDED

#include "api/thingspect_event.pb.h"

#include <functional>
#include <grpcpp/generic/async_generic_service.h>
#include <grpcpp/support/async_stream.h>
#include <grpcpp/support/async_unary_call.h>
#include <grpcpp/support/client_callback.h>
#include <grpcpp/client_context.h>
#include <grpcpp/completion_queue.h>
#include <grpcpp/support/message_allocator.h>
#include <grpcpp/support/method_handler.h>
#include <grpcpp/impl/proto_utils.h>
#include <grpcpp/impl/rpc_method.h>
#include <grpcpp/support/server_callback.h>
#include <grpcpp/impl/server_callback_handlers.h>
#include <grpcpp/server_context.h>
#include <grpcpp/impl/service_type.h>
#include <grpcpp/support/status.h>
#include <grpcpp/support/stub_options.h>
#include <grpcpp/support/sync_stream.h>

namespace thingspect {
namespace api {

// EventService contains functions to query events.
class EventService final {
 public:
  static constexpr char const* service_full_name() {
    return "thingspect.api.EventService";
  }
  class StubInterface {
   public:
    virtual ~StubInterface() {}
    // List all events for a device in an [end, start) time range, in descending timestamp order. Events are generated by rules when conditions are met.
    virtual ::grpc::Status ListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::thingspect::api::ListEventsResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::ListEventsResponse>> AsyncListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::ListEventsResponse>>(AsyncListEventsRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::ListEventsResponse>> PrepareAsyncListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::ListEventsResponse>>(PrepareAsyncListEventsRaw(context, request, cq));
    }
    // List the latest event for each of an organization's devices. Events are generated by rules when conditions are met.
    virtual ::grpc::Status LatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::thingspect::api::LatestEventsResponse* response) = 0;
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::LatestEventsResponse>> AsyncLatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::LatestEventsResponse>>(AsyncLatestEventsRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::LatestEventsResponse>> PrepareAsyncLatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::LatestEventsResponse>>(PrepareAsyncLatestEventsRaw(context, request, cq));
    }
    class async_interface {
     public:
      virtual ~async_interface() {}
      // List all events for a device in an [end, start) time range, in descending timestamp order. Events are generated by rules when conditions are met.
      virtual void ListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest* request, ::thingspect::api::ListEventsResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void ListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest* request, ::thingspect::api::ListEventsResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
      // List the latest event for each of an organization's devices. Events are generated by rules when conditions are met.
      virtual void LatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest* request, ::thingspect::api::LatestEventsResponse* response, std::function<void(::grpc::Status)>) = 0;
      virtual void LatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest* request, ::thingspect::api::LatestEventsResponse* response, ::grpc::ClientUnaryReactor* reactor) = 0;
    };
    typedef class async_interface experimental_async_interface;
    virtual class async_interface* async() { return nullptr; }
    class async_interface* experimental_async() { return async(); }
   private:
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::ListEventsResponse>* AsyncListEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::ListEventsResponse>* PrepareAsyncListEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::LatestEventsResponse>* AsyncLatestEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) = 0;
    virtual ::grpc::ClientAsyncResponseReaderInterface< ::thingspect::api::LatestEventsResponse>* PrepareAsyncLatestEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) = 0;
  };
  class Stub final : public StubInterface {
   public:
    Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());
    ::grpc::Status ListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::thingspect::api::ListEventsResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::ListEventsResponse>> AsyncListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::ListEventsResponse>>(AsyncListEventsRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::ListEventsResponse>> PrepareAsyncListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::ListEventsResponse>>(PrepareAsyncListEventsRaw(context, request, cq));
    }
    ::grpc::Status LatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::thingspect::api::LatestEventsResponse* response) override;
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::LatestEventsResponse>> AsyncLatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::LatestEventsResponse>>(AsyncLatestEventsRaw(context, request, cq));
    }
    std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::LatestEventsResponse>> PrepareAsyncLatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) {
      return std::unique_ptr< ::grpc::ClientAsyncResponseReader< ::thingspect::api::LatestEventsResponse>>(PrepareAsyncLatestEventsRaw(context, request, cq));
    }
    class async final :
      public StubInterface::async_interface {
     public:
      void ListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest* request, ::thingspect::api::ListEventsResponse* response, std::function<void(::grpc::Status)>) override;
      void ListEvents(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest* request, ::thingspect::api::ListEventsResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
      void LatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest* request, ::thingspect::api::LatestEventsResponse* response, std::function<void(::grpc::Status)>) override;
      void LatestEvents(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest* request, ::thingspect::api::LatestEventsResponse* response, ::grpc::ClientUnaryReactor* reactor) override;
     private:
      friend class Stub;
      explicit async(Stub* stub): stub_(stub) { }
      Stub* stub() { return stub_; }
      Stub* stub_;
    };
    class async* async() override { return &async_stub_; }

   private:
    std::shared_ptr< ::grpc::ChannelInterface> channel_;
    class async async_stub_{this};
    ::grpc::ClientAsyncResponseReader< ::thingspect::api::ListEventsResponse>* AsyncListEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::thingspect::api::ListEventsResponse>* PrepareAsyncListEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::ListEventsRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::thingspect::api::LatestEventsResponse>* AsyncLatestEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) override;
    ::grpc::ClientAsyncResponseReader< ::thingspect::api::LatestEventsResponse>* PrepareAsyncLatestEventsRaw(::grpc::ClientContext* context, const ::thingspect::api::LatestEventsRequest& request, ::grpc::CompletionQueue* cq) override;
    const ::grpc::internal::RpcMethod rpcmethod_ListEvents_;
    const ::grpc::internal::RpcMethod rpcmethod_LatestEvents_;
  };
  static std::unique_ptr<Stub> NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options = ::grpc::StubOptions());

  class Service : public ::grpc::Service {
   public:
    Service();
    virtual ~Service();
    // List all events for a device in an [end, start) time range, in descending timestamp order. Events are generated by rules when conditions are met.
    virtual ::grpc::Status ListEvents(::grpc::ServerContext* context, const ::thingspect::api::ListEventsRequest* request, ::thingspect::api::ListEventsResponse* response);
    // List the latest event for each of an organization's devices. Events are generated by rules when conditions are met.
    virtual ::grpc::Status LatestEvents(::grpc::ServerContext* context, const ::thingspect::api::LatestEventsRequest* request, ::thingspect::api::LatestEventsResponse* response);
  };
  template <class BaseClass>
  class WithAsyncMethod_ListEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_ListEvents() {
      ::grpc::Service::MarkMethodAsync(0);
    }
    ~WithAsyncMethod_ListEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ListEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestListEvents(::grpc::ServerContext* context, ::thingspect::api::ListEventsRequest* request, ::grpc::ServerAsyncResponseWriter< ::thingspect::api::ListEventsResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithAsyncMethod_LatestEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithAsyncMethod_LatestEvents() {
      ::grpc::Service::MarkMethodAsync(1);
    }
    ~WithAsyncMethod_LatestEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LatestEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestLatestEvents(::grpc::ServerContext* context, ::thingspect::api::LatestEventsRequest* request, ::grpc::ServerAsyncResponseWriter< ::thingspect::api::LatestEventsResponse>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  typedef WithAsyncMethod_ListEvents<WithAsyncMethod_LatestEvents<Service > > AsyncService;
  template <class BaseClass>
  class WithCallbackMethod_ListEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_ListEvents() {
      ::grpc::Service::MarkMethodCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::thingspect::api::ListEventsRequest, ::thingspect::api::ListEventsResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::thingspect::api::ListEventsRequest* request, ::thingspect::api::ListEventsResponse* response) { return this->ListEvents(context, request, response); }));}
    void SetMessageAllocatorFor_ListEvents(
        ::grpc::MessageAllocator< ::thingspect::api::ListEventsRequest, ::thingspect::api::ListEventsResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(0);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::thingspect::api::ListEventsRequest, ::thingspect::api::ListEventsResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_ListEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ListEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* ListEvents(
      ::grpc::CallbackServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithCallbackMethod_LatestEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithCallbackMethod_LatestEvents() {
      ::grpc::Service::MarkMethodCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::thingspect::api::LatestEventsRequest, ::thingspect::api::LatestEventsResponse>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::thingspect::api::LatestEventsRequest* request, ::thingspect::api::LatestEventsResponse* response) { return this->LatestEvents(context, request, response); }));}
    void SetMessageAllocatorFor_LatestEvents(
        ::grpc::MessageAllocator< ::thingspect::api::LatestEventsRequest, ::thingspect::api::LatestEventsResponse>* allocator) {
      ::grpc::internal::MethodHandler* const handler = ::grpc::Service::GetHandler(1);
      static_cast<::grpc::internal::CallbackUnaryHandler< ::thingspect::api::LatestEventsRequest, ::thingspect::api::LatestEventsResponse>*>(handler)
              ->SetMessageAllocator(allocator);
    }
    ~WithCallbackMethod_LatestEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LatestEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* LatestEvents(
      ::grpc::CallbackServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/)  { return nullptr; }
  };
  typedef WithCallbackMethod_ListEvents<WithCallbackMethod_LatestEvents<Service > > CallbackService;
  typedef CallbackService ExperimentalCallbackService;
  template <class BaseClass>
  class WithGenericMethod_ListEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_ListEvents() {
      ::grpc::Service::MarkMethodGeneric(0);
    }
    ~WithGenericMethod_ListEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ListEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithGenericMethod_LatestEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithGenericMethod_LatestEvents() {
      ::grpc::Service::MarkMethodGeneric(1);
    }
    ~WithGenericMethod_LatestEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LatestEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
  };
  template <class BaseClass>
  class WithRawMethod_ListEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_ListEvents() {
      ::grpc::Service::MarkMethodRaw(0);
    }
    ~WithRawMethod_ListEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ListEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestListEvents(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(0, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawMethod_LatestEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawMethod_LatestEvents() {
      ::grpc::Service::MarkMethodRaw(1);
    }
    ~WithRawMethod_LatestEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LatestEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    void RequestLatestEvents(::grpc::ServerContext* context, ::grpc::ByteBuffer* request, ::grpc::ServerAsyncResponseWriter< ::grpc::ByteBuffer>* response, ::grpc::CompletionQueue* new_call_cq, ::grpc::ServerCompletionQueue* notification_cq, void *tag) {
      ::grpc::Service::RequestAsyncUnary(1, context, request, response, new_call_cq, notification_cq, tag);
    }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_ListEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_ListEvents() {
      ::grpc::Service::MarkMethodRawCallback(0,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->ListEvents(context, request, response); }));
    }
    ~WithRawCallbackMethod_ListEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status ListEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* ListEvents(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithRawCallbackMethod_LatestEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithRawCallbackMethod_LatestEvents() {
      ::grpc::Service::MarkMethodRawCallback(1,
          new ::grpc::internal::CallbackUnaryHandler< ::grpc::ByteBuffer, ::grpc::ByteBuffer>(
            [this](
                   ::grpc::CallbackServerContext* context, const ::grpc::ByteBuffer* request, ::grpc::ByteBuffer* response) { return this->LatestEvents(context, request, response); }));
    }
    ~WithRawCallbackMethod_LatestEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable synchronous version of this method
    ::grpc::Status LatestEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    virtual ::grpc::ServerUnaryReactor* LatestEvents(
      ::grpc::CallbackServerContext* /*context*/, const ::grpc::ByteBuffer* /*request*/, ::grpc::ByteBuffer* /*response*/)  { return nullptr; }
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_ListEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_ListEvents() {
      ::grpc::Service::MarkMethodStreamed(0,
        new ::grpc::internal::StreamedUnaryHandler<
          ::thingspect::api::ListEventsRequest, ::thingspect::api::ListEventsResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::thingspect::api::ListEventsRequest, ::thingspect::api::ListEventsResponse>* streamer) {
                       return this->StreamedListEvents(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_ListEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status ListEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::ListEventsRequest* /*request*/, ::thingspect::api::ListEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedListEvents(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::thingspect::api::ListEventsRequest,::thingspect::api::ListEventsResponse>* server_unary_streamer) = 0;
  };
  template <class BaseClass>
  class WithStreamedUnaryMethod_LatestEvents : public BaseClass {
   private:
    void BaseClassMustBeDerivedFromService(const Service* /*service*/) {}
   public:
    WithStreamedUnaryMethod_LatestEvents() {
      ::grpc::Service::MarkMethodStreamed(1,
        new ::grpc::internal::StreamedUnaryHandler<
          ::thingspect::api::LatestEventsRequest, ::thingspect::api::LatestEventsResponse>(
            [this](::grpc::ServerContext* context,
                   ::grpc::ServerUnaryStreamer<
                     ::thingspect::api::LatestEventsRequest, ::thingspect::api::LatestEventsResponse>* streamer) {
                       return this->StreamedLatestEvents(context,
                         streamer);
                  }));
    }
    ~WithStreamedUnaryMethod_LatestEvents() override {
      BaseClassMustBeDerivedFromService(this);
    }
    // disable regular version of this method
    ::grpc::Status LatestEvents(::grpc::ServerContext* /*context*/, const ::thingspect::api::LatestEventsRequest* /*request*/, ::thingspect::api::LatestEventsResponse* /*response*/) override {
      abort();
      return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
    }
    // replace default version of method with streamed unary
    virtual ::grpc::Status StreamedLatestEvents(::grpc::ServerContext* context, ::grpc::ServerUnaryStreamer< ::thingspect::api::LatestEventsRequest,::thingspect::api::LatestEventsResponse>* server_unary_streamer) = 0;
  };
  typedef WithStreamedUnaryMethod_ListEvents<WithStreamedUnaryMethod_LatestEvents<Service > > StreamedUnaryService;
  typedef Service SplitStreamedService;
  typedef WithStreamedUnaryMethod_ListEvents<WithStreamedUnaryMethod_LatestEvents<Service > > StreamedService;
};

}  // namespace api
}  // namespace thingspect


#endif  // GRPC_api_2fthingspect_5fevent_2eproto__INCLUDED
