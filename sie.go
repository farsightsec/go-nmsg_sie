package nmsg_sie

import "github.com/farsightsec/go-nmsg"

//go:generate ./genproto.sh

func (p *Delay) GetVid() uint32      { return 2 }
func (p *DnsDedupe) GetVid() uint32  { return 2 }
func (p *DnsNX) GetVid() uint32      { return 2 }
func (p *NewDomain) GetVid() uint32  { return 2 }
func (p *QR) GetVid() uint32         { return 2 }
func (p *Reputation) GetVid() uint32 { return 2 }

func (p *Delay) GetMsgtype() uint32      { return 4 }
func (p *DnsDedupe) GetMsgtype() uint32  { return 1 }
func (p *DnsNX) GetMsgtype() uint32      { return 6 }
func (p *NewDomain) GetMsgtype() uint32  { return 5 }
func (p *QR) GetMsgtype() uint32         { return 3 }
func (p *Reputation) GetMsgtype() uint32 { return 2 }

func init() {
	nmsg.Register(&Delay{})
	nmsg.Register(&DnsDedupe{})
	nmsg.Register(&DnsNX{})
	nmsg.Register(&NewDomain{})
	nmsg.Register(&QR{})
	nmsg.Register(&Reputation{})
}
