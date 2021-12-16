package valhalla

import (
	"fmt"
	"testing"
)

func TestElevationPolyline(t *testing.T) {
	t.Skip()
	elevationRequest := ElevationRequest{}
	elevationRequest.EncodedPolyline = "xxvmg@frkwoHuC`@qFxAwEfB_Bx@o@`BQpB\\pDtAjElB~DlAzDjAvFt@nGOlBy@lAe{fQsNlFq}tCsXb@oHNiHj@eC`AoEpAaI@wJe@iKeEmIoJ?bHc@|NGxBu@|D_AnDqFtM{FfM_LvV}A~DaApEgLva@EjA{BtFwAfFoFnRyIf{gAdDoNnj@{d@zmAwBjDqFrHuSpU_W~W_LnJyI|JcCtD{C`GgEpNiNl^sFjKiYri@eMfU{KrUoFfLaH~QaJj_@uKdi@aUzuAy@jF}AfKqA`JgAvK{@dJu@~Kg@lMkAnZoB`a@wB`l@cCrn@qBzr@KnNc@pVKfT\\|PdAtQ|@xKnAbLhB`MpBhMvBnL|Fn`@bAdEjB|P|Jlh@lBhKxQlfAl@hErAlEbQ~_@hJnTtEvKzCxItC`JtBpJvAnJt@fNZrMx@n{jA|MhF~^f@vF`CjUlD|VjGbd@nBhJzBdJtFxU~BdPvDvk@pAzE`CxQhApE~DjXfD~StCdQzFpT|EnPpDbXvDk@vSha@pPt^`BbDj@`Ec@dBcA`B{AlAwc@hR{@rA_@xBHzBd@nBzo@|`AtTp\\x`@`g@bZl_@tQ|{nC~EjBnBb@Rr@HlAEf@G`@Kr@Sn@_@l@y@x@kA`CyFpBsEbCqF~@oAb@Yl@g@fA?zAHjAFzGdAfI|@zFbB~DbDv\\hX~HfItDdEjHjMdEzH|@vBPfCFfDFvAb@xCpAxAxCxCp@TjNrKxBrAfB`AvAj@xAp@vAVjCPzCDfEGdRGnEXtEr@xDrApE~BhOlK`GjFpIdKjIbJhEpClEdBjDfAlJdCbMtCjHp@fK^zGPpKM`FEr_AkGre@eU`t@ea@~HwDdJaD`GeAtFcA|`@LlG@dLz@fSpAhT~BlEQjDa@~JsA~OuBpS}@lA@`ZNtEEvJB|EAnIL~Pv@~IfBvXxD`JdAvMfBdRb@bLLfFWrKStI}@vHwAf\\{HjPoBnU_Dt|@gMxJs@bR@fNZpW`@pe@bB~CVlDf@rFhAxCt@bG|AzMpEhL|BjX|Hjb@nTlK|DzK~ExM`H~J|DbOpBtHp@nKrA`NIpI_@`Gm@xGaA`FiAfIaAtFg@~GOrp@O_Ae^s@yMyAuK"
	elevationRequest.Range = true
	elevations, err := client.Height(elevationRequest)
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}
	fmt.Printf("Elevations: %+v", elevations)
}

func TestElevationShape(t *testing.T) {
	elevationRequest := ElevationRequest{}
	elevationRequest.Shape = []ElevationPoint{{
		Lat: -41.2924,
		Lon: 174.7787,
	},
		{
			Lat: -36.8509,
			Lon: 174.7645,
		}}
	elevationRequest.Range = true
	elevations, err := client.Height(elevationRequest)
	if err != nil {
		t.Fatalf("Error: %+v", err)
	}
	if len(elevations.RangeHeight) != len(elevationRequest.Shape) {
		t.Errorf("Expected %d results got %d", len(elevationRequest.Shape), len(elevations.RangeHeight))
	}

}
