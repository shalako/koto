// Code generated by protoc-gen-twirp v5.12.0, DO NOT EDIT.
// source: info.proto

package rpc

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// =====================
// InfoService Interface
// =====================

type InfoService interface {
	PublicKey(context.Context, *Empty) (*InfoPublicKeyResponse, error)

	Version(context.Context, *Empty) (*InfoVersionResponse, error)
}

// ===========================
// InfoService Protobuf Client
// ===========================

type infoServiceProtobufClient struct {
	client HTTPClient
	urls   [2]string
	opts   twirp.ClientOptions
}

// NewInfoServiceProtobufClient creates a Protobuf client that implements the InfoService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewInfoServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) InfoService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + InfoServicePathPrefix
	urls := [2]string{
		prefix + "PublicKey",
		prefix + "Version",
	}

	return &infoServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *infoServiceProtobufClient) PublicKey(ctx context.Context, in *Empty) (*InfoPublicKeyResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	out := new(InfoPublicKeyResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *infoServiceProtobufClient) Version(ctx context.Context, in *Empty) (*InfoVersionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	out := new(InfoVersionResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// =======================
// InfoService JSON Client
// =======================

type infoServiceJSONClient struct {
	client HTTPClient
	urls   [2]string
	opts   twirp.ClientOptions
}

// NewInfoServiceJSONClient creates a JSON client that implements the InfoService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewInfoServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) InfoService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + InfoServicePathPrefix
	urls := [2]string{
		prefix + "PublicKey",
		prefix + "Version",
	}

	return &infoServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *infoServiceJSONClient) PublicKey(ctx context.Context, in *Empty) (*InfoPublicKeyResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	out := new(InfoPublicKeyResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *infoServiceJSONClient) Version(ctx context.Context, in *Empty) (*InfoVersionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	out := new(InfoVersionResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==========================
// InfoService Server Handler
// ==========================

type infoServiceServer struct {
	InfoService
	hooks *twirp.ServerHooks
}

func NewInfoServiceServer(svc InfoService, hooks *twirp.ServerHooks) TwirpServer {
	return &infoServiceServer{
		InfoService: svc,
		hooks:       hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *infoServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// InfoServicePathPrefix is used for all URL paths on a twirp InfoService server.
// Requests are always: POST InfoServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const InfoServicePathPrefix = "/rpc.InfoService/"

func (s *infoServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/rpc.InfoService/PublicKey":
		s.servePublicKey(ctx, resp, req)
		return
	case "/rpc.InfoService/Version":
		s.serveVersion(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *infoServiceServer) servePublicKey(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePublicKeyJSON(ctx, resp, req)
	case "application/protobuf":
		s.servePublicKeyProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *infoServiceServer) servePublicKeyJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *InfoPublicKeyResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InfoService.PublicKey(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoPublicKeyResponse and nil error while calling PublicKey. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) servePublicKeyProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *InfoPublicKeyResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InfoService.PublicKey(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoPublicKeyResponse and nil error while calling PublicKey. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) serveVersion(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveVersionJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveVersionProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *infoServiceServer) serveVersionJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *InfoVersionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InfoService.Version(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoVersionResponse and nil error while calling Version. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) serveVersionProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *InfoVersionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InfoService.Version(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoVersionResponse and nil error while calling Version. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor2, 0
}

func (s *infoServiceServer) ProtocGenTwirpVersion() string {
	return "v5.12.0"
}

func (s *infoServiceServer) PathPrefix() string {
	return InfoServicePathPrefix
}

var twirpFileDescriptor2 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcc, 0x4b, 0xcb,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0xce, 0xcd, 0x4f,
	0x49, 0xcd, 0x81, 0x88, 0x28, 0x99, 0x71, 0x89, 0x7a, 0xe6, 0xa5, 0xe5, 0x07, 0x94, 0x26, 0xe5,
	0x64, 0x26, 0x7b, 0xa7, 0x56, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x72,
	0x71, 0x15, 0x80, 0x05, 0xe3, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38,
	0x0b, 0x60, 0xca, 0x94, 0x4c, 0xb8, 0x84, 0x41, 0xfa, 0xc2, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3,
	0x90, 0x75, 0xa5, 0xe4, 0x27, 0x67, 0xa7, 0x16, 0xc5, 0x97, 0x24, 0xa6, 0xc3, 0x74, 0x41, 0x44,
	0x42, 0x12, 0xd3, 0x8d, 0x8a, 0xb9, 0xb8, 0x41, 0xba, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53,
	0x85, 0x8c, 0xb9, 0x38, 0xe1, 0x16, 0x0b, 0x71, 0xe9, 0x15, 0x15, 0x24, 0xeb, 0xb9, 0xe6, 0x16,
	0x94, 0x54, 0x4a, 0x49, 0x81, 0xd9, 0xd8, 0x1d, 0xa6, 0xcf, 0xc5, 0x0e, 0xb5, 0x15, 0x45, 0x8b,
	0x04, 0x5c, 0x0b, 0x9a, 0x9b, 0x9c, 0x38, 0xa2, 0xd8, 0xf4, 0xf4, 0xf4, 0x8b, 0x0a, 0x92, 0x93,
	0xd8, 0xc0, 0x7e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x3c, 0x3b, 0x29, 0x13, 0x01,
	0x00, 0x00,
}
