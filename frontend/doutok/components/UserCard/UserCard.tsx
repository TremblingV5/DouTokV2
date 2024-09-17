"use client";

import { Button, Card, Divider, message } from "antd";
import Avatar from "antd/es/avatar/avatar";
import Meta from "antd/es/card/Meta";
import { useState } from "react";

import {
  UserServiceGetUserInfoResponse,
  useUserServiceGetUserInfo
} from "@/api/svapi/api";
import "./UserCard.css";
import { LoginModal } from "@/components/LoginModal/LoginModal";
import useUserStore from "@/components/UserStore/useUserStore";
import { UpdateUserInfoForm } from "@/components/UpdateUserInfoForm/UpdateUserInfoForm";

const failedGetUserInfo = () => {
  message.error("获取用户信息失败，请重新登录");
};

export function UserCard() {
  const [openLoginModal, setOpenLoginModal] = useState<boolean>(false);

  const removeToken = useUserStore(state => state.removeToken);

  const [avatar, setAvatar] = useState<string | undefined>("no-login.svg");
  const [username, setUsername] = useState<string | undefined>();
  const [following, setFollowing] = useState<string | undefined>();
  const [fans, setFans] = useState<string | undefined>();
  const [likes, setLikes] = useState<string | undefined>();
  const [doutokId, setDouTokId] = useState<string | undefined>();

  useUserServiceGetUserInfo({
    resolve: (resp: UserServiceGetUserInfoResponse) => {
      const { data } = resp;
      if (resp.code !== 0 || data === undefined) {
        failedGetUserInfo();
        return resp;
      }

      setAvatar(data.user?.avatar);
      setUsername(data.user?.name);
      setFollowing(data.user?.followCount ? data.user?.followCount : "0");
      setFans(data.user?.followerCount ? data.user?.followerCount : "0");
      setDouTokId(data.user?.id);
      setLikes("0");
      return resp;
    }
  });

  const [oepnEditUserInfoModal, setOpenEditUserInfoModal] =
    useState<boolean>(false);

  return (
    <div className={"user-card-container"}>
      <Card>
        <Meta
          avatar={
            <Avatar
              src={
                avatar != undefined && avatar.length != 0
                  ? avatar
                  : "no-login.svg"
              }
              size={150}
            />
          }
          title={
            <span className={"username"}>
              {username != undefined ? username : "未登录"}
            </span>
          }
          description={
            <>
              <div className={"user-stats"}>
                <div className={"following-num"}>
                  <span>关注: {following != undefined ? following : "-"}</span>
                </div>
                <Divider className={"divider"} type={"horizontal"} />
                <div className={"fans-num"}>
                  <span>粉丝: {fans != undefined ? following : "-"}</span>
                </div>
                <Divider className={"divider"} type={"horizontal"} />
                <div className={"likes-num"}>
                  <span>获赞: {likes != undefined ? following : "-"}</span>
                </div>
              </div>
            </>
          }
        />
        {doutokId != undefined && (
          <>
            <div className={"buttons"}>
              <div className={"doutok-id"}>
                DouTok号: {doutokId != undefined ? doutokId : "-"}
              </div>
              <div className={"logout-button"}>
                <Button
                  onClick={() => {
                    removeToken();
                    localStorage.removeItem("token");
                    window.location.reload();
                  }}
                >
                  退出登陆
                </Button>
              </div>
              <div className={"edit-button"}>
                <Button onClick={() => setOpenEditUserInfoModal(true)}>
                  编辑资料
                </Button>
              </div>
            </div>
          </>
        )}
        {doutokId == undefined && (
          <>
            <div className={"buttons"}>
              <div className={"edit-button"}>
                <Button onClick={() => setOpenLoginModal(true)}>
                  登录DouTok
                </Button>
              </div>
            </div>
          </>
        )}
      </Card>
      {openLoginModal && (
        <LoginModal
          open={openLoginModal}
          onCancel={() => setOpenLoginModal(false)}
          type={"login"}
        />
      )}
      {oepnEditUserInfoModal && (
        <UpdateUserInfoForm
          open={oepnEditUserInfoModal}
          onCancel={() => setOpenEditUserInfoModal(false)}
        />
      )}
    </div>
  );
}
