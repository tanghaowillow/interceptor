package flexfec

import "github.com/pion/interceptor"

type FecDecodingInterceptor struct {
	// BindRTCPReader lets you modify any incoming RTCP packets. It is called once per sender/receiver, however this might
	// change in the future. The returned method will be called once per packet batch.

}

// BindRTCPReader lets you modify any incoming RTCP packets. It is called once per sender/receiver, however this might
// change in the future. The returned method will be called once per packet batch.
// BindRTCPReader(reader RTCPReader) RTCPReader

// func (r *FecDecodingInterceptor) BindRTCPReader(reader interceptor.RTCPReader) interceptor.RTCPReader {

// }

// BindRemoteStream lets you modify any incoming RTP packets. It is called once for per RemoteStream. The returned method
// will be called once per rtp packet.
// BindRemoteStream(info *StreamInfo, reader RTPReader) RTPReader

func (r *FecDecodingInterceptor)