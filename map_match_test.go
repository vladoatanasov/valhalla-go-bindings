package valhalla

import (
	"encoding/json"
	"testing"
)

func TestTraceRoute(t *testing.T) {
	request := TraceRouteRequest{}
	request.Costing = "auto"
	request.EncodedPolyline = "u}vfaBzcxF^|CRjD?bAStBo@dC}@nBkAtAo@j@eD^gAKeA_@}@q@o@{@c@{@a@eASw@U_AWaBQmBuC_Ns@_D]uAwBuFa@iAcAaA}@_@w@IoB?yABwQrEiGbAkCzAm@\\}ALaAEOCy@Mw@Yy@_@qAo@iB[cCe@uBiAuBFwE\\yCf@eEX]@iATyA^mBf@oCbBs@bAw@hAs@zA_AtB_AxBmDjIgKxOSNaG|EkJ~EeG|BmDp@qFv@mANgHL{IP_DMs@CaHOmBCqAAoDGaFIiBCYAcEIuDGuAAuFPkCNi@DgCf@cE~@wBt@}At@{OrIuGbDqAp@y@`@sDfBoB`AuAp@oGbDqAp@aChAwFxC{FpCo@\\wAj@wAr@YLoAj@uItE{A`Co@XcDtAuCfAs@Gu@o@sAXwJ|CyC|@iK|CoVbIaA\\sAd@kEzAgBh@kA^eA\\iFbBgBf@{Bt@uTvHsWtB{AD{ANyIt@GcFMsGUaEQiC[qEq@sJ_AwKwAaOg@wGe@oFk@{Gq@sG}B}Vm@aHuAsP_AaNSsCgBuVMiB[sDk@qFcCnB}@r@{AjAcD`D_JrHeRxMQLeAv@oClBiCpB{AbAm@b@cAt@iG`FaDrC_D`D_CpCmEhFkC~CkFfGcUzVgAhAcVbXoKvLwBzBy@z@{@z@}AdB{BdC_@`@s@v@gRtSiIxIu@t@k@l@sBbC{AbBs@v@_AdA}@`AwP`Su\\p^kB|By@~@i@l@yAdB}AbBg@p@{AfBwEhFkg@vj@oCdDmAzAg@l@yBbCcCnCeAfAw@|@OP_DpDegAtlAcAjAm@t@}B~BeBdBa@b@u@t@sMnNkBtBo@r@]`@sGrHa@`@_@^wAxAz@|ChBhGxHxWZdA\\dA~BvH~@hDbBrF|FvShIfZvDfNx@zCd@fBhBzGaFdDaAr@oAv@gK~Gm@\\uG`Eg@XaDzAeJnEyDnBoE|BwAx@}A~@qGpDkDlBsI~FiAbFmCrAgBx@kBx@sC{AyLbFw@TeCh@aB\\sB\\kBPcBB_A@}QRw@Bu@@cDDmCFcA?_A?gN?iPL_A@i`@JiCBiBMoFCgAFqDCoBA{CCoA@{A?_BE{IDyBCyB?kB?mFD{C@cB@_C?qC?aGCaAtA{F@eKWq@qAeSd@kPn@yGVmAhCm@?iABkGh@sGdAmInAiP|De[vGeARsGtAuGnAyDr@}JtB_MfCkCZoABa@GmBq@kc@kB{WcA_g@}AkIeAs@ScDuA_CkAaBo@oAWiBU_E_@}CHw@HoB\\{NjF{FtDmHnGiRfReAhAsDxDq@r@}@~@gm@jo@gSlTsAxAgDbDmE~D_Ax@am@`g@URw@l@cBrAeDnCaI`H{CrBsM|H}@vDaDvBgAl@gAp@}BnAwDfCel@|e@yHxG}KrJeTpPwHjIyE|GmEdHwRb[gE~Gi@`A{F~JSZy@|AsAbCsJvQkAvBe@nAa@|@eDnGeB|Aw@pBwLbYqBpHuI`WmDfL_CvIcAbEgAfEo@tCeCxK}AzH}AfKoB`OuBjPcBhO]|BYpBgAxH}@`H[|B[lCi@rEg@nEwEbb@}ArNiAlJ{CnT}A`JkBvJo@fDq@~CiBxHgDzNwF|RTrBYzAG\\eCvJaAfDoCfL^xID|CAbEg@pCVzh@Nb_@BdFD~G@nCC~AF|^j@bcAVpf@\\nh@NxQf@dm@^li@^re@LrP`Alc@f@~N|@pVj@tJp@nMdC|_@p@pMrCve@pDrk@h@|JtAnTrB`]lBf[|Crd@x@fN|@tMNzBLhB~AlVLtB|@fNtCrc@~AtVzDxi@RbD`Dnd@tAnQpApHxEjTpCvHjD|JdKpU~B|J~@zG`@dFzAbY|AxZtA`RHxAf@bGLhCZtE\\rF\\xIg@`FRlCv@bOz@vNNxDx@|Qb@rJ_@tGm@|BkA`A_Ax@cATaBf@aE`@qZoDuSeAcGD{IH_R|BmAFqCNcBRgB\\cDz@gC|@}ClAwGbDuGxDuIdG{D~CyDxDkLrNgEjH{BtEwAhD}BjGmCnIoSfu@{BtIsEnPuCdJmAlDeDbIyCpE_EdGyFbHyDzE{@|@aBbBqQ|Q_F`FoHhG_CzA}CpBuG|DmE~Cgh@ne@sQbP}C|CwE`FmCvCyA|AcBxAeA`AoAz@}AfAqFvC{VdM}HdDeEnBgCxAwChDaBxA_D`DeCxCaQ~RsQl^eFdJSb@gA|BaEbIoLfUmD|GmFzJqBpDsBtDyRn_@sBtDmN~VsMpTwCxFyFpJyEhLaDfJgEdOcB`JuA|LuBnV}@jJ_@hFmGno@qDxP[nAmEbSqDnHcR|_@kd@taAuF~L}@pBcE~IeF`Lw@fBmC`GkYpn@eNjZsFtLqOp]cJlScMpV{JpPcJpNyDrF{IhL}G`IcJhJmGlF}FhEmGhDcHbDcFrBcMzEuQxGoIdDmG~CoF~CkFnD_HrFsGlGoF~EyDjD_DrCaF~CeAvAmAbAcAf@{A`Au@t@i@`Aa@pAQpAGbBHtSCzGI|GYxH}@zMqCjWoDbUkl@xkCeK`n@uK`d@{G~OyG~OwFzKmGvJ{@bAqDrEmNbKsg@dZklB`z@ayCtnAsAn@iHfDqe@hUsC|AsmApo@wAnAqhAz|@uf@~c@{IdJcb@~b@_[l^qQxSoPrRuaBrnBuT~YuWv]yInLmStg@wGhPsYlx@wb@lhA{Lj^a_@xgAeZv|@}G~RoJx\\qXx`Aud@hpBeEfNsMfa@aShe@mOvVmMjOmJdIsYhT}S`Mog@l\\}KpGgA\\cFbBuHr@}QKqF~@eEr@u@FoCd@gFjAwFpCyGjEgy@z`AqSbW{Up^wQdZoO|VmExH{a@dt@sHvH}FjH{NvQeKbVqEpKkExL_DpLkCrLoBrKoIjk@_BjJoAvGkA|E_ArEeEnNgDtHgDhF_EfH}FdHaGzF}FhEoEhCeG`CaGbB}FhA_Fx@wIOmUeD}H{Bov@oRk{Ac_@aHeBmDkBeWcHcOqDuFcA_LeCgh@eMiA[gZ}GoRuCq`@iCc[i@wg@{@amAqB}pAeCwSWsWDkWfB{UjF{OdGe[pPy@h@aHjHoUnR}GvKkOnV{A`F_ElJcEnKyTrh@iBnD{Rj_@_Tfa@}NfSobA`fAmr@jm@_d@jWePhG_a@fM{QxCml@vF_OvAsNrAspAfTgq@~Ju^rFad@hGgWlDyrAlDoFTsPn@ob@fAeVlCkQtFcNjHwK~HeMdLiIlKuRnZoG~M_E|LsEvNeBnHmAbEs@jCMdBq@vPWjEi@`Dq@|BaBdD_ChC_EnCoGpIuHdQuElNon@|bBcRpg@ai@tvA}O~`@yKbWmSh^aQvU}Wfd@uMv\\mQnh@aB|FuAjFiEpReFdZoFzZuC`TU`BIbDR~Cz@`Ed@|DJhEQnEg@lDk@dBi@`BoAxB{BpBaF`AoC_@uCkByCkFyA}HmEuIsDiEgE}EoIqEuCoAgCc@mB]iIg@qLkAa{BkEknBmF_KOcJU{h@yAwXm@kLUs|BiGgq@_Gqq@eKgk@mJyk@cFe\\Iw[`BeVvB_PvCya@xJu`@|Ocz@`c@y}@f`@wvAvg@yLlDuRpHyExA}JjBmJ~AyJXiUImW{EmJyB}PsGcNiHuLoHy_@oVcA_@}CUoDIwCjBkD~@qC]iCkAsB_BuA}AgDyHoBaG_BkFs@eEa@sD{C_GyA_CaeCshBaSgNsRyLuTaMiTsKso@{Vcn@uP}l@eKg{AmTycBwUywBo[ii@yHg}@_McY}DykFmu@kZoEyXuFaZqIsXaL}VaNkPyKmNgKsMqK{wO__NmV_TqKuIgKsH_O{Jca@yVerLw|G}[eSewGywDoTgMqSqMkfAeo@um@{b@ap@eg@aX_Tsq@kh@c\\qUea@}Zc|@ib@yLgEgNeF}i@mSeS{FcV{Fwn@wJu`@yFm^yA{p@oEyW{BcoA_FkUkCg^kFgs@cNmx@cQqy@cU{p@cUwt@s^uz@ce@q`A{i@w_@_RyYoK{\\qIuTaEu_@}Fqm@uAud@zAui@lFu_@`Hme@fO}]tOqYxO}^b\\aj@rc@wbA~q@ih@h\\}[pNsm@`Pcv@jNoo@vAiSDkw@uEqgAkGowAkF{i@kAov@dEel@pImcAdTw`Avf@ov@dn@{p@tu@{eDjfEacDh`EkxA|kBq_AfdAm]r_@kYbU_}@do@{WhRgb@~Tea@hRusDhuAc~CbqAg{BrcAecDvyAazAhq@oa@`Q_`@dLcl@vLuo@|Gyr@fCkg@m@{`@qCgY{D{SsDyRmEaT_GgSsG}RoGqEiBmLaFeQ{HkOqHgk@aYeHyEoViPya@a`@kVsXgP}Q{LwMaWmXqVoYiZi`@kQiWef@uv@sUwZm[a^eT{Ss[yWocBixAm^eWoZsPyb@aTmg@eWgd@sSi_@cRuYyRuIeGeZ_Tc_@sZgWyTm`@ca@iZ_\\uOqR}Vq[{Ws_@yP}W}OcW_e@gu@sd@as@q_@co@kRsZ{QcYuNoSs[y`@uRsSeTyRkPcN}LcJ{N{JuWsO}ViMke@oPoMeEyM{Cc]mG}RcCmM}@ka@wAoREkNRs^bCmOtAsNtBmPbDmVxFaVlI{_@vOeMhGoOlIwM|I}J|GsJjH_JtHmTnQcLzIsN|KgKzHsP`LmSrL{YxN}[|LqR`Gq\\pHcLtBcLzAeObBwNlAe`@jAcHAgI?sV{@cUsAuXeD}\\uG_c@}J{NyCiQoCmSyBmN{@sW]}RHyLZ{NbAkMfAsZfFiTrEsMpC{MnC{YzFi^tHq]hHaTxEcRlDu]nHa^fHef@fKe^lHqRxDmX~Fcf@rJwuAzYg_@zGe_@bE_^hCiOh@cS\\{SFeS_@qN]gSuA_TkB}\\sEwf@_Jez@kQcf@kLudAeZs~@mXuh@eRgTkK{RqKcSaNuQwMgRgPyXqYgO{QcNuQiVa^kLgRoJuOoWqd@kXue@ki@k_A_q@ulAe\\mk@opAe}BufA}lByKeSwKsQoIqNgPoXeUg]qE}GoRyV{NsQoKkLyJyKq[wZ{UmSe_@_YuWqPkd@oVma@eQkSaHcQiFsPmEaXeGi`@cHgS{CoVyD{r@qKeX{Du^_G{a@{G}i@yHg]iH{q@oPsYqIs\\{Kst@sYkkA_m@oIeEu^ySwl@y^go@yb@ugCwdB{m@eb@s^cWeq@yd@}cBqjA_|AwdAizAacAmg@o]sc@eYo[qQmi@sWca@_Pgf@qPwj@mNcg@mJgb@{Fwg@}Dgg@qBmq@OsX^kXjA{UjB_[~@q`ExRgb@vBgb@nBudAdF}u@zDwsArGgq@jDcq@hDsW|BcWzC}RvCcXjFa[vHsf@|M_a@`Mk`@bNmnAta@if@dPqf@`Pg[`LmVbIcj@zQep@nTij@zQq_@lMe[nKon@lS{s@hV}n@xSge@vO_d@dPuZ|Jkh@|Pww@nW_x@hXg}@hZuj@xQg\\dLul@jTgZrIqQjGak@vRca@jPw`@dQk[lO{_@xSyc@|Wcw@hk@y_@nW_[~WaY~Um]h\\}k@vm@sj@lp@mb@fi@{f@`n@ceA|sAyzArmBaf@tm@ki@lq@}|@`iAsR~Via@dh@uZt`@}b@`j@az@heAa^fe@e^rc@qg@fm@ae@je@}[lZeY|Uaf@x]oRhM_PpJoh@tXc[pNuRxHkZzKeW~HiW|GuRfEeRzD}`@zGkWxCsb@zD_p@|Bsz@Ecz@yDwo@uHik@yJao@_PwReGw^gMkXqLkZeNsUwL{V}NaU{Ns^iVut@cg@qh@w]w}@cn@smBgqAka@iXer@w`@wv@ob@ie@_Tei@cTaXkKuQcGuUeIgZuKul@uS{k@iSsoAcc@ceAm^cbAq^wo@kTil@_Tq]wL{g@mQoaAa]ol@{Si|@{Zmx@sXmV}Ikb@{Oeg@uPwg@mQ_h@mPifA_Ywc@iJic@wHub@wGaTeCiTeCcc@iEkn@eEqh@kDyWiBwe@gCkk@cDoe@{CoUyAgVwA_`@qBwSeBkW{Bek@{Hg\\qFm\\yGwe@yLqWaIy[iK_ReHoRuH}VuKc`@cRwg@sYyj@e\\w_@wYcj@sb@wp@_n@ka@wa@c\\{^qb@sh@a}@cnAef@cs@w~@kqAsg@it@aTsYgTqX_NePsL{MyZyZqVaUwj@yd@e]iU_h@_[ql@uXac@aQw]uK_TwFaYkGmT_E}N}B{NiBm^aDi_@aCw^a@{d@HqYp@_TzAeTzAiT`Cut@`Ia_AzJqTnCsXxCum@tG{o@jHap@fH{l@dHsn@rG{`AvKe}@dJk}AvPed@hFwiA|Lml@fD_j@`B_n@d@ke@y@m]w@uv@cEyqAeIwa@uCmb@sCscAoG_x@_F_y@}EeWsA{Wc@uUGwYj@}TbA{Ip@eNpAkU~BsX~EsUnEcYnHaTxGuWpJ{UlJcVnLgv@lb@_f@hXofAfn@at@fa@ol@t\\oYhPmy@he@ie@pWae@dXmw@rc@{q@p`@wQ|KyYlOi[bRgL`H{_Azg@am@tW_S`I}VpIa]bK}_@bKq^hGmh@zI_c@tJiXtH}f@xPcn@bYm_@lTgh@~]ybAns@yh@l_@ym@te@}y@`q@ie@z^oOfLo@b@kAx@sg@d\\{`Ajh@iDtAwQ`Io{A~s@kBj@{OrIuLjGqf@rY_j@|]_P|KsNlIwLbH{D`CgCzAgGrDiEhCqWjMui@~Xgp@p^mm@n]wcAbl@o}@ji@uu@de@uh@l[k{@ff@o]pSi]lQ{`Avg@{v@`c@u`A`k@{_Ahj@erAdx@oKpGsbApl@_d@`WgkAtp@co@b^ql@|]cc@`Xs\\|RmBlAm^fUaj@bc@op@tf@qh@b_@_h@z[o{BxrAsiA|o@u}@zg@cy@pd@cEzBuiAbn@qt@lb@op@p\\ix@p`@m|@de@um@d\\gl@nYmo@d]sUlLaU~Kgh@tXqQdKu^xUgKfG}IzHwUlTmSnS}HlGwIzFcGlFoFhFkEzFkCrIkCbKjAnI`@nI[lJqBbIoBxDkD`BsEAaGmBuGPiHbEoKtJoO|UoPvVoN|S{LfSeYph@{LlTgJjQaX`k@uThf@aJpTaSjg@aOpe@uMzg@qLzg@eLfk@uIfk@mEx[}CtWoBvRqChYqBlTuB~Wy@hMcCl`@kAxRkAhVy@vVgBng@eFzpAw@|VuEn}A{Adf@KjHuD|}@yEp_BwEbrAsA~WyB~WaAvNuBbSqCnV}DxY{Hdc@kJz^{Kpb@sJvYcKbU_LdWoPxY{_@zf@_MtM_QpQeXvVcQjNqQbNyU`OmS`L_MlG_ZjNwYlKeZjJmZvIuVtFsGfAyLlBq^`Eod@tB{c@Ta^k@{c@_AwLQgYc@gp@?"
	r, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	t.Logf("Request: %s", r)
	resp, err := client.TraceRoute(request)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	t.Logf("Result: %+v", resp)
}
func TestTraceAttributes(t *testing.T) {
	request := TraceAttributesRequest{}
	request.Costing = "auto"
	request.EncodedPolyline = "u}vfaBzcxF^|CRjD?bAStBo@dC}@nBkAtAo@j@eD^gAKeA_@}@q@o@{@c@{@a@eASw@U_AWaBQmBuC_Ns@_D]uAwBuFa@iAcAaA}@_@w@IoB?yABwQrEiGbAkCzAm@\\}ALaAEOCy@Mw@Yy@_@qAo@iB[cCe@uBiAuBFwE\\yCf@eEX]@iATyA^mBf@oCbBs@bAw@hAs@zA_AtB_AxBmDjIgKxOSNaG|EkJ~EeG|BmDp@qFv@mANgHL{IP_DMs@CaHOmBCqAAoDGaFIiBCYAcEIuDGuAAuFPkCNi@DgCf@cE~@wBt@}At@{OrIuGbDqAp@y@`@sDfBoB`AuAp@oGbDqAp@aChAwFxC{FpCo@\\wAj@wAr@YLoAj@uItE{A`Co@XcDtAuCfAs@Gu@o@sAXwJ|CyC|@iK|CoVbIaA\\sAd@kEzAgBh@kA^eA\\iFbBgBf@{Bt@uTvHsWtB{AD{ANyIt@GcFMsGUaEQiC[qEq@sJ_AwKwAaOg@wGe@oFk@{Gq@sG}B}Vm@aHuAsP_AaNSsCgBuVMiB[sDk@qFcCnB}@r@{AjAcD`D_JrHeRxMQLeAv@oClBiCpB{AbAm@b@cAt@iG`FaDrC_D`D_CpCmEhFkC~CkFfGcUzVgAhAcVbXoKvLwBzBy@z@{@z@}AdB{BdC_@`@s@v@gRtSiIxIu@t@k@l@sBbC{AbBs@v@_AdA}@`AwP`Su\\p^kB|By@~@i@l@yAdB}AbBg@p@{AfBwEhFkg@vj@oCdDmAzAg@l@yBbCcCnCeAfAw@|@OP_DpDegAtlAcAjAm@t@}B~BeBdBa@b@u@t@sMnNkBtBo@r@]`@sGrHa@`@_@^wAxAz@|ChBhGxHxWZdA\\dA~BvH~@hDbBrF|FvShIfZvDfNx@zCd@fBhBzGaFdDaAr@oAv@gK~Gm@\\uG`Eg@XaDzAeJnEyDnBoE|BwAx@}A~@qGpDkDlBsI~FiAbFmCrAgBx@kBx@sC{AyLbFw@TeCh@aB\\sB\\kBPcBB_A@}QRw@Bu@@cDDmCFcA?_A?gN?iPL_A@i`@JiCBiBMoFCgAFqDCoBA{CCoA@{A?_BE{IDyBCyB?kB?mFD{C@cB@_C?qC?aGCaAtA{F@eKWq@qAeSd@kPn@yGVmAhCm@?iABkGh@sGdAmInAiP|De[vGeARsGtAuGnAyDr@}JtB_MfCkCZoABa@GmBq@kc@kB{WcA_g@}AkIeAs@ScDuA_CkAaBo@oAWiBU_E_@}CHw@HoB\\{NjF{FtDmHnGiRfReAhAsDxDq@r@}@~@gm@jo@gSlTsAxAgDbDmE~D_Ax@am@`g@URw@l@cBrAeDnCaI`H{CrBsM|H}@vDaDvBgAl@gAp@}BnAwDfCel@|e@yHxG}KrJeTpPwHjIyE|GmEdHwRb[gE~Gi@`A{F~JSZy@|AsAbCsJvQkAvBe@nAa@|@eDnGeB|Aw@pBwLbYqBpHuI`WmDfL_CvIcAbEgAfEo@tCeCxK}AzH}AfKoB`OuBjPcBhO]|BYpBgAxH}@`H[|B[lCi@rEg@nEwEbb@}ArNiAlJ{CnT}A`JkBvJo@fDq@~CiBxHgDzNwF|RTrBYzAG\\eCvJaAfDoCfL^xID|CAbEg@pCVzh@Nb_@BdFD~G@nCC~AF|^j@bcAVpf@\\nh@NxQf@dm@^li@^re@LrP`Alc@f@~N|@pVj@tJp@nMdC|_@p@pMrCve@pDrk@h@|JtAnTrB`]lBf[|Crd@x@fN|@tMNzBLhB~AlVLtB|@fNtCrc@~AtVzDxi@RbD`Dnd@tAnQpApHxEjTpCvHjD|JdKpU~B|J~@zG`@dFzAbY|AxZtA`RHxAf@bGLhCZtE\\rF\\xIg@`FRlCv@bOz@vNNxDx@|Qb@rJ_@tGm@|BkA`A_Ax@cATaBf@aE`@qZoDuSeAcGD{IH_R|BmAFqCNcBRgB\\cDz@gC|@}ClAwGbDuGxDuIdG{D~CyDxDkLrNgEjH{BtEwAhD}BjGmCnIoSfu@{BtIsEnPuCdJmAlDeDbIyCpE_EdGyFbHyDzE{@|@aBbBqQ|Q_F`FoHhG_CzA}CpBuG|DmE~Cgh@ne@sQbP}C|CwE`FmCvCyA|AcBxAeA`AoAz@}AfAqFvC{VdM}HdDeEnBgCxAwChDaBxA_D`DeCxCaQ~RsQl^eFdJSb@gA|BaEbIoLfUmD|GmFzJqBpDsBtDyRn_@sBtDmN~VsMpTwCxFyFpJyEhLaDfJgEdOcB`JuA|LuBnV}@jJ_@hFmGno@qDxP[nAmEbSqDnHcR|_@kd@taAuF~L}@pBcE~IeF`Lw@fBmC`GkYpn@eNjZsFtLqOp]cJlScMpV{JpPcJpNyDrF{IhL}G`IcJhJmGlF}FhEmGhDcHbDcFrBcMzEuQxGoIdDmG~CoF~CkFnD_HrFsGlGoF~EyDjD_DrCaF~CeAvAmAbAcAf@{A`Au@t@i@`Aa@pAQpAGbBHtSCzGI|GYxH}@zMqCjWoDbUkl@xkCeK`n@uK`d@{G~OyG~OwFzKmGvJ{@bAqDrEmNbKsg@dZklB`z@ayCtnAsAn@iHfDqe@hUsC|AsmApo@wAnAqhAz|@uf@~c@{IdJcb@~b@_[l^qQxSoPrRuaBrnBuT~YuWv]yInLmStg@wGhPsYlx@wb@lhA{Lj^a_@xgAeZv|@}G~RoJx\\qXx`Aud@hpBeEfNsMfa@aShe@mOvVmMjOmJdIsYhT}S`Mog@l\\}KpGgA\\cFbBuHr@}QKqF~@eEr@u@FoCd@gFjAwFpCyGjEgy@z`AqSbW{Up^wQdZoO|VmExH{a@dt@sHvH}FjH{NvQeKbVqEpKkExL_DpLkCrLoBrKoIjk@_BjJoAvGkA|E_ArEeEnNgDtHgDhF_EfH}FdHaGzF}FhEoEhCeG`CaGbB}FhA_Fx@wIOmUeD}H{Bov@oRk{Ac_@aHeBmDkBeWcHcOqDuFcA_LeCgh@eMiA[gZ}GoRuCq`@iCc[i@wg@{@amAqB}pAeCwSWsWDkWfB{UjF{OdGe[pPy@h@aHjHoUnR}GvKkOnV{A`F_ElJcEnKyTrh@iBnD{Rj_@_Tfa@}NfSobA`fAmr@jm@_d@jWePhG_a@fM{QxCml@vF_OvAsNrAspAfTgq@~Ju^rFad@hGgWlDyrAlDoFTsPn@ob@fAeVlCkQtFcNjHwK~HeMdLiIlKuRnZoG~M_E|LsEvNeBnHmAbEs@jCMdBq@vPWjEi@`Dq@|BaBdD_ChC_EnCoGpIuHdQuElNon@|bBcRpg@ai@tvA}O~`@yKbWmSh^aQvU}Wfd@uMv\\mQnh@aB|FuAjFiEpReFdZoFzZuC`TU`BIbDR~Cz@`Ed@|DJhEQnEg@lDk@dBi@`BoAxB{BpBaF`AoC_@uCkByCkFyA}HmEuIsDiEgE}EoIqEuCoAgCc@mB]iIg@qLkAa{BkEknBmF_KOcJU{h@yAwXm@kLUs|BiGgq@_Gqq@eKgk@mJyk@cFe\\Iw[`BeVvB_PvCya@xJu`@|Ocz@`c@y}@f`@wvAvg@yLlDuRpHyExA}JjBmJ~AyJXiUImW{EmJyB}PsGcNiHuLoHy_@oVcA_@}CUoDIwCjBkD~@qC]iCkAsB_BuA}AgDyHoBaG_BkFs@eEa@sD{C_GyA_CaeCshBaSgNsRyLuTaMiTsKso@{Vcn@uP}l@eKg{AmTycBwUywBo[ii@yHg}@_McY}DykFmu@kZoEyXuFaZqIsXaL}VaNkPyKmNgKsMqK{wO__NmV_TqKuIgKsH_O{Jca@yVerLw|G}[eSewGywDoTgMqSqMkfAeo@um@{b@ap@eg@aX_Tsq@kh@c\\qUea@}Zc|@ib@yLgEgNeF}i@mSeS{FcV{Fwn@wJu`@yFm^yA{p@oEyW{BcoA_FkUkCg^kFgs@cNmx@cQqy@cU{p@cUwt@s^uz@ce@q`A{i@w_@_RyYoK{\\qIuTaEu_@}Fqm@uAud@zAui@lFu_@`Hme@fO}]tOqYxO}^b\\aj@rc@wbA~q@ih@h\\}[pNsm@`Pcv@jNoo@vAiSDkw@uEqgAkGowAkF{i@kAov@dEel@pImcAdTw`Avf@ov@dn@{p@tu@{eDjfEacDh`EkxA|kBq_AfdAm]r_@kYbU_}@do@{WhRgb@~Tea@hRusDhuAc~CbqAg{BrcAecDvyAazAhq@oa@`Q_`@dLcl@vLuo@|Gyr@fCkg@m@{`@qCgY{D{SsDyRmEaT_GgSsG}RoGqEiBmLaFeQ{HkOqHgk@aYeHyEoViPya@a`@kVsXgP}Q{LwMaWmXqVoYiZi`@kQiWef@uv@sUwZm[a^eT{Ss[yWocBixAm^eWoZsPyb@aTmg@eWgd@sSi_@cRuYyRuIeGeZ_Tc_@sZgWyTm`@ca@iZ_\\uOqR}Vq[{Ws_@yP}W}OcW_e@gu@sd@as@q_@co@kRsZ{QcYuNoSs[y`@uRsSeTyRkPcN}LcJ{N{JuWsO}ViMke@oPoMeEyM{Cc]mG}RcCmM}@ka@wAoREkNRs^bCmOtAsNtBmPbDmVxFaVlI{_@vOeMhGoOlIwM|I}J|GsJjH_JtHmTnQcLzIsN|KgKzHsP`LmSrL{YxN}[|LqR`Gq\\pHcLtBcLzAeObBwNlAe`@jAcHAgI?sV{@cUsAuXeD}\\uG_c@}J{NyCiQoCmSyBmN{@sW]}RHyLZ{NbAkMfAsZfFiTrEsMpC{MnC{YzFi^tHq]hHaTxEcRlDu]nHa^fHef@fKe^lHqRxDmX~Fcf@rJwuAzYg_@zGe_@bE_^hCiOh@cS\\{SFeS_@qN]gSuA_TkB}\\sEwf@_Jez@kQcf@kLudAeZs~@mXuh@eRgTkK{RqKcSaNuQwMgRgPyXqYgO{QcNuQiVa^kLgRoJuOoWqd@kXue@ki@k_A_q@ulAe\\mk@opAe}BufA}lByKeSwKsQoIqNgPoXeUg]qE}GoRyV{NsQoKkLyJyKq[wZ{UmSe_@_YuWqPkd@oVma@eQkSaHcQiFsPmEaXeGi`@cHgS{CoVyD{r@qKeX{Du^_G{a@{G}i@yHg]iH{q@oPsYqIs\\{Kst@sYkkA_m@oIeEu^ySwl@y^go@yb@ugCwdB{m@eb@s^cWeq@yd@}cBqjA_|AwdAizAacAmg@o]sc@eYo[qQmi@sWca@_Pgf@qPwj@mNcg@mJgb@{Fwg@}Dgg@qBmq@OsX^kXjA{UjB_[~@q`ExRgb@vBgb@nBudAdF}u@zDwsArGgq@jDcq@hDsW|BcWzC}RvCcXjFa[vHsf@|M_a@`Mk`@bNmnAta@if@dPqf@`Pg[`LmVbIcj@zQep@nTij@zQq_@lMe[nKon@lS{s@hV}n@xSge@vO_d@dPuZ|Jkh@|Pww@nW_x@hXg}@hZuj@xQg\\dLul@jTgZrIqQjGak@vRca@jPw`@dQk[lO{_@xSyc@|Wcw@hk@y_@nW_[~WaY~Um]h\\}k@vm@sj@lp@mb@fi@{f@`n@ceA|sAyzArmBaf@tm@ki@lq@}|@`iAsR~Via@dh@uZt`@}b@`j@az@heAa^fe@e^rc@qg@fm@ae@je@}[lZeY|Uaf@x]oRhM_PpJoh@tXc[pNuRxHkZzKeW~HiW|GuRfEeRzD}`@zGkWxCsb@zD_p@|Bsz@Ecz@yDwo@uHik@yJao@_PwReGw^gMkXqLkZeNsUwL{V}NaU{Ns^iVut@cg@qh@w]w}@cn@smBgqAka@iXer@w`@wv@ob@ie@_Tei@cTaXkKuQcGuUeIgZuKul@uS{k@iSsoAcc@ceAm^cbAq^wo@kTil@_Tq]wL{g@mQoaAa]ol@{Si|@{Zmx@sXmV}Ikb@{Oeg@uPwg@mQ_h@mPifA_Ywc@iJic@wHub@wGaTeCiTeCcc@iEkn@eEqh@kDyWiBwe@gCkk@cDoe@{CoUyAgVwA_`@qBwSeBkW{Bek@{Hg\\qFm\\yGwe@yLqWaIy[iK_ReHoRuH}VuKc`@cRwg@sYyj@e\\w_@wYcj@sb@wp@_n@ka@wa@c\\{^qb@sh@a}@cnAef@cs@w~@kqAsg@it@aTsYgTqX_NePsL{MyZyZqVaUwj@yd@e]iU_h@_[ql@uXac@aQw]uK_TwFaYkGmT_E}N}B{NiBm^aDi_@aCw^a@{d@HqYp@_TzAeTzAiT`Cut@`Ia_AzJqTnCsXxCum@tG{o@jHap@fH{l@dHsn@rG{`AvKe}@dJk}AvPed@hFwiA|Lml@fD_j@`B_n@d@ke@y@m]w@uv@cEyqAeIwa@uCmb@sCscAoG_x@_F_y@}EeWsA{Wc@uUGwYj@}TbA{Ip@eNpAkU~BsX~EsUnEcYnHaTxGuWpJ{UlJcVnLgv@lb@_f@hXofAfn@at@fa@ol@t\\oYhPmy@he@ie@pWae@dXmw@rc@{q@p`@wQ|KyYlOi[bRgL`H{_Azg@am@tW_S`I}VpIa]bK}_@bKq^hGmh@zI_c@tJiXtH}f@xPcn@bYm_@lTgh@~]ybAns@yh@l_@ym@te@}y@`q@ie@z^oOfLo@b@kAx@sg@d\\{`Ajh@iDtAwQ`Io{A~s@kBj@{OrIuLjGqf@rY_j@|]_P|KsNlIwLbH{D`CgCzAgGrDiEhCqWjMui@~Xgp@p^mm@n]wcAbl@o}@ji@uu@de@uh@l[k{@ff@o]pSi]lQ{`Avg@{v@`c@u`A`k@{_Ahj@erAdx@oKpGsbApl@_d@`WgkAtp@co@b^ql@|]cc@`Xs\\|RmBlAm^fUaj@bc@op@tf@qh@b_@_h@z[o{BxrAsiA|o@u}@zg@cy@pd@cEzBuiAbn@qt@lb@op@p\\ix@p`@m|@de@um@d\\gl@nYmo@d]sUlLaU~Kgh@tXqQdKu^xUgKfG}IzHwUlTmSnS}HlGwIzFcGlFoFhFkEzFkCrIkCbKjAnI`@nI[lJqBbIoBxDkD`BsEAaGmBuGPiHbEoKtJoO|UoPvVoN|S{LfSeYph@{LlTgJjQaX`k@uThf@aJpTaSjg@aOpe@uMzg@qLzg@eLfk@uIfk@mEx[}CtWoBvRqChYqBlTuB~Wy@hMcCl`@kAxRkAhVy@vVgBng@eFzpAw@|VuEn}A{Adf@KjHuD|}@yEp_BwEbrAsA~WyB~WaAvNuBbSqCnV}DxY{Hdc@kJz^{Kpb@sJvYcKbU_LdWoPxY{_@zf@_MtM_QpQeXvVcQjNqQbNyU`OmS`L_MlG_ZjNwYlKeZjJmZvIuVtFsGfAyLlBq^`Eod@tB{c@Ta^k@{c@_AwLQgYc@gp@?"
	r, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	t.Logf("Request: %s", r)
	resp, err := client.TraceAttributes(request)
	if err != nil {
		t.Fatalf("Error: %s", err.Error())
	}
	t.Logf("Result: %+v", resp)
}
