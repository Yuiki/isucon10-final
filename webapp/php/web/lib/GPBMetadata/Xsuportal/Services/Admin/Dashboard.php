<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/services/admin/dashboard.proto

namespace GPBMetadata\Xsuportal\Services\Admin;

class Dashboard
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Xsuportal\Resources\Leaderboard::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0a89020a28787375706f7274616c2f73657276696365732f61646d696e2f" .
            "64617368626f6172642e70726f746f121e787375706f7274616c2e70726f" .
            "746f2e73657276696365732e61646d696e22120a1044617368626f617264" .
            "5265717565737422500a1144617368626f617264526573706f6e7365123b" .
            "0a0b6c6561646572626f61726418012001280b32262e787375706f727461" .
            "6c2e70726f746f2e7265736f75726365732e4c6561646572626f61726442" .
            "4f5a4d6769746875622e636f6d2f697375636f6e2f697375636f6e31302d" .
            "66696e616c2f7765626170702f676f6c616e672f70726f746f2f78737570" .
            "6f7274616c2f73657276696365732f61646d696e620670726f746f33"
        ));

        static::$is_initialized = true;
    }
}
