<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xsuportal/resources/benchmark_result.proto

namespace Xsuportal\Proto\Resources;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xsuportal.proto.resources.BenchmarkResult</code>
 */
class BenchmarkResult extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bool finished = 1;</code>
     */
    private $finished = false;
    /**
     * Generated from protobuf field <code>bool passed = 2;</code>
     */
    private $passed = false;
    /**
     * Generated from protobuf field <code>int64 score = 3;</code>
     */
    private $score = 0;
    /**
     * Generated from protobuf field <code>.xsuportal.proto.resources.BenchmarkResult.ScoreBreakdown score_breakdown = 4;</code>
     */
    private $score_breakdown = null;
    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     */
    private $reason = '';
    /**
     * string stdout = 6;
     * string stderr = 7;
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp marked_at = 6;</code>
     */
    private $marked_at = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type bool $finished
     *     @type bool $passed
     *     @type int|string $score
     *     @type \Xsuportal\Proto\Resources\BenchmarkResult\ScoreBreakdown $score_breakdown
     *     @type string $reason
     *     @type \Google\Protobuf\Timestamp $marked_at
     *           string stdout = 6;
     *           string stderr = 7;
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Xsuportal\Resources\BenchmarkResult::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bool finished = 1;</code>
     * @return bool
     */
    public function getFinished()
    {
        return $this->finished;
    }

    /**
     * Generated from protobuf field <code>bool finished = 1;</code>
     * @param bool $var
     * @return $this
     */
    public function setFinished($var)
    {
        GPBUtil::checkBool($var);
        $this->finished = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool passed = 2;</code>
     * @return bool
     */
    public function getPassed()
    {
        return $this->passed;
    }

    /**
     * Generated from protobuf field <code>bool passed = 2;</code>
     * @param bool $var
     * @return $this
     */
    public function setPassed($var)
    {
        GPBUtil::checkBool($var);
        $this->passed = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int64 score = 3;</code>
     * @return int|string
     */
    public function getScore()
    {
        return $this->score;
    }

    /**
     * Generated from protobuf field <code>int64 score = 3;</code>
     * @param int|string $var
     * @return $this
     */
    public function setScore($var)
    {
        GPBUtil::checkInt64($var);
        $this->score = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xsuportal.proto.resources.BenchmarkResult.ScoreBreakdown score_breakdown = 4;</code>
     * @return \Xsuportal\Proto\Resources\BenchmarkResult\ScoreBreakdown
     */
    public function getScoreBreakdown()
    {
        return $this->score_breakdown;
    }

    /**
     * Generated from protobuf field <code>.xsuportal.proto.resources.BenchmarkResult.ScoreBreakdown score_breakdown = 4;</code>
     * @param \Xsuportal\Proto\Resources\BenchmarkResult\ScoreBreakdown $var
     * @return $this
     */
    public function setScoreBreakdown($var)
    {
        GPBUtil::checkMessage($var, \Xsuportal\Proto\Resources\BenchmarkResult_ScoreBreakdown::class);
        $this->score_breakdown = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     * @return string
     */
    public function getReason()
    {
        return $this->reason;
    }

    /**
     * Generated from protobuf field <code>string reason = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->reason = $var;

        return $this;
    }

    /**
     * string stdout = 6;
     * string stderr = 7;
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp marked_at = 6;</code>
     * @return \Google\Protobuf\Timestamp
     */
    public function getMarkedAt()
    {
        return $this->marked_at;
    }

    /**
     * string stdout = 6;
     * string stderr = 7;
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp marked_at = 6;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setMarkedAt($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->marked_at = $var;

        return $this;
    }

}
