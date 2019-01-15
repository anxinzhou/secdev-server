package contract

// Get All functions in contract

type AssembleContract struct {
	PbcERC998 *PbcERC998
	PbcERC20 *Pbc
	PvcERC998 *PvcERC998
	PvcERC20 *Pvc
}

func NewAssembleContract (pbcERC998 *PbcERC998,pbcERC20 *Pbc,pvcERC998 *PvcERC998,pvcERC20 *Pvc) *AssembleContract {
	return &AssembleContract{
		PbcERC998:pbcERC998,
		PbcERC20:pbcERC20,
		PvcERC998:pvcERC998,
		PvcERC20:pvcERC20,
	}
}
