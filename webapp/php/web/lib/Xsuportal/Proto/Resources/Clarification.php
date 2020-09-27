<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/resources/clarification.proto

namespace Xsuportal\Proto\Resources;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xsuportal.proto.resources.Clarification</code>
 */
class Clarification extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     */
    private $id = 0;
    /**
     * Generated from protobuf field <code>int64 team_id = 2;</code>
     */
    private $team_id = 0;
    /**
     * Generated from protobuf field <code>bool answered = 3;</code>
     */
    private $answered = false;
    /**
     * Generated from protobuf field <code>bool disclosed = 4;</code>
     */
    private $disclosed = false;
    /**
     * Generated from protobuf field <code>string question = 5;</code>
     */
    private $question = '';
    /**
     * Generated from protobuf field <code>string answer = 6;</code>
     */
    private $answer = '';
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp created_at = 7;</code>
     */
    private $created_at = null;
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp answered_at = 8;</code>
     */
    private $answered_at = null;
    /**
     * Generated from protobuf field <code>.xsuportal.proto.resources.Team team = 16;</code>
     */
    private $team = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int|string $id
     *     @type int|string $team_id
     *     @type bool $answered
     *     @type bool $disclosed
     *     @type string $question
     *     @type string $answer
     *     @type \Google\Protobuf\Timestamp $created_at
     *     @type \Google\Protobuf\Timestamp $answered_at
     *     @type \Xsuportal\Proto\Resources\Team $team
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Xsuportal\Resources\Clarification::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @return int|string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>int64 id = 1;</code>
     * @param int|string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkInt64($var);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 team_id = 2;</code>
     * @return int|string
     */
    public function getTeamId()
    {
        return $this->team_id;
    }

    /**
     * Generated from protobuf field <code>int64 team_id = 2;</code>
     * @param int|string $var
     * @return $this
     */
    public function setTeamId($var)
    {
        GPBUtil::checkInt64($var);
        $this->team_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool answered = 3;</code>
     * @return bool
     */
    public function getAnswered()
    {
        return $this->answered;
    }

    /**
     * Generated from protobuf field <code>bool answered = 3;</code>
     * @param bool $var
     * @return $this
     */
    public function setAnswered($var)
    {
        GPBUtil::checkBool($var);
        $this->answered = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool disclosed = 4;</code>
     * @return bool
     */
    public function getDisclosed()
    {
        return $this->disclosed;
    }

    /**
     * Generated from protobuf field <code>bool disclosed = 4;</code>
     * @param bool $var
     * @return $this
     */
    public function setDisclosed($var)
    {
        GPBUtil::checkBool($var);
        $this->disclosed = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string question = 5;</code>
     * @return string
     */
    public function getQuestion()
    {
        return $this->question;
    }

    /**
     * Generated from protobuf field <code>string question = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setQuestion($var)
    {
        GPBUtil::checkString($var, True);
        $this->question = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string answer = 6;</code>
     * @return string
     */
    public function getAnswer()
    {
        return $this->answer;
    }

    /**
     * Generated from protobuf field <code>string answer = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setAnswer($var)
    {
        GPBUtil::checkString($var, True);
        $this->answer = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp created_at = 7;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getCreatedAt()
    {
        return $this->created_at;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp created_at = 7;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setCreatedAt($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->created_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp answered_at = 8;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getAnsweredAt()
    {
        return $this->answered_at;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp answered_at = 8;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setAnsweredAt($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->answered_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xsuportal.proto.resources.Team team = 16;</code>
     * @return \Xsuportal\Proto\Resources\Team
     */
    public function getTeam()
    {
        return $this->team;
    }

    /**
     * Generated from protobuf field <code>.xsuportal.proto.resources.Team team = 16;</code>
     * @param \Xsuportal\Proto\Resources\Team $var
     * @return $this
     */
    public function setTeam($var)
    {
        GPBUtil::checkMessage($var, \Xsuportal\Proto\Resources\Team::class);
        $this->team = $var;

        return $this;
    }

}
