-- доделать
ALTER FUNCTION [dbo].[UspGetMemberUser](@Login NVARCHAR(50))
RETURNS TABLE
AS
RETURN
(
    SELECT
        -- TblUser
        u.mRegDate AS User_mRegDate,
        u.mUserAuth,
        u.mUserNo AS User_mUserNo,
        u.mUserId AS User_mUserId,
        u.mUserPswd AS User_mUserPswd,
        u.mCertifiedKey,
        u.mIp AS User_mIp,
        u.mLoginTm AS User_mLoginTm,
        u.mLogoutTm,
        u.mTotUseTm,
        u.mWorldNo,
        u.mDelDate AS User_mDelDate,
        u.mPcBangLv,
        u.mSecKeyTableUse,
        u.mUseMacro,
        u.mIpEX,
        u.mJoinCode,
        u.mLoginChannelID,
        u.mTired,
        u.mChnSID,
        u.mNewId,
        u.mLoginSvrType,
        u.mAccountGuid,
        u.mNormalLimitTime,
        u.mPcBangLimitTime,
        u.mRegIp,
        u.mIsMovingToBattleSvr,

        -- Member
        m.mUserId AS Member_mUserId,
        m.mUserPswd AS Member_mUserPswd,
        m.Superpwd,
        m.Cash,
        m.email AS Member_Email,
        m.tgzh,
        m.uid,
        m.klq,
        m.ylq,
        m.auth,
        m.mSum,
        m.isadmin,
        m.isdl,
        m.dlmoney,
        m.registerIp,
        m.country,
        m.CashBack,

        -- TblPc
        p.mRegDate AS Pc_mRegDate,
        p.mOwner,
        p.mSlot,
        p.mNo AS Pc_mNo,
        p.mNm,
        p.mClass,
        p.mSex,
        p.mHead,
        p.mFace,
        p.mBody,
        p.mHomeMapNo,
        p.mHomePosX,
        p.mHomePosY,
        p.mHomePosZ,
        p.mDelDate AS Pc_mDelDate,

        -- TblPcState
        pt.mNo AS PcState_mNo,
        pt.mLevel,
        pt.mExp,
        pt.mHpAdd,
        pt.mHp,
        pt.mMpAdd,
        pt.mMp,
        pt.mMapNo,
        pt.mPosX,
        pt.mPosY,
        pt.mPosZ,
        pt.mStomach,
        pt.mIp AS PcState_mIp,
        pt.mLoginTm AS PcState_mLoginTm,
        pt.mLogoutTm AS PcState_mLogoutTm,
        pt.mTotUseTm AS PcState_mTotUseTm,
        pt.mPkCnt,
        pt.mChaotic,
        pt.mDiscipleJoinCount,
        pt.mPartyMemCntLevel,
        pt.mLostExp,
        pt.mIsLetterLimit,
        pt.mFlag,
        pt.mIsPreventItemDrop,
        pt.mSkillTreePoint,
        pt.mRestExpGuild,
        pt.mRestExpActivate,
        pt.mRestExpDeactivate,
        pt.mQMCnt,
        pt.mGuildQMCnt,
        pt.mFierceCnt,
        pt.mBossCnt,

        -- TblPcInventory
        pi.mRegDate AS PcInv_mRegDate,
        pi.mSerialNo AS PcInv_mSerialNo,
        pi.mPcNo AS PcInv_mPcNo,
        pi.mItemNo AS PcInv_mItemNo,
        pi.mEndDate AS PcInv_mEndDate,
        pi.mIsConfirm AS PcInv_mIsConfirm,
        pi.mStatus AS PcInv_mStatus,
        pi.mCnt AS PcInv_mCnt,
        pi.mCntUse AS PcInv_mCntUse,
        pi.mIsSeizure AS PcInv_mIsSeizure,
        pi.mApplyAbnItemNo AS PcInv_mApplyAbnItemNo,
        pi.mApplyAbnItemEndDate AS PcInv_mApplyAbnItemEndDate,
        pi.mOwner AS PcInv_mOwner,
        pi.mPracticalPeriod AS PcInv_mPracticalPeriod,
        pi.mBindingType AS PcInv_mBindingType,
        pi.mRestoreCnt AS PcInv_mRestoreCnt,
        pi.mHoleCount AS PcInv_mHoleCount,

        -- TblPcStore
        ps.mRegDate AS PcStore_mRegDate,
        ps.mSerialNo AS PcStore_mSerialNo,
        ps.mUserNo AS PcStore_mUserNo,
        ps.mItemNo AS PcStore_mItemNo,
        ps.mEndDate AS PcStore_mEndDate,
        ps.mIsConfirm AS PcStore_mIsConfirm,
        ps.mStatus AS PcStore_mStatus,
        ps.mCnt AS PcStore_mCnt,
        ps.mCntUse AS PcStore_mCntUse,
        ps.mIsSeizure AS PcStore_mIsSeizure,
        ps.mApplyAbnItemNo AS PcStore_mApplyAbnItemNo,
        ps.mApplyAbnItemEndDate AS PcStore_mApplyAbnItemEndDate,
        ps.mOwner AS PcStore_mOwner,
        ps.mPracticalPeriod AS PcStore_mPracticalPeriod,
        ps.mBindingType AS PcStore_mBindingType,
        ps.mRestoreCnt AS PcStore_mRestoreCnt,
        ps.mHoleCount AS PcStore_mHoleCount
    FROM
        FNLAccount.[dbo].[TblUser] u
    INNER JOIN
        FNLAccount.[dbo].[Member] m ON u.mUserId = m.mUserId
    INNER JOIN
        FNLGame2155.[dbo].[TblPc] p ON p.mOwner = u.mUserNo
    LEFT JOIN
        FNLGame2155.[dbo].[TblPcState] pt ON pt.mNo = p.mNo
    LEFT JOIN
        FNLGame2155.[dbo].[TblPcInventory] pi ON pi.mPcNo = p.mNo
    LEFT JOIN
        FNLGame2155.[dbo].[TblPcStore] ps ON ps.mUserNo = p.mOwner
    WHERE
        m.mUserId = @Login OR m.email = @Login
)