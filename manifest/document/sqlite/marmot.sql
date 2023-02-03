DROP TABLE IF EXISTS inspection;
CREATE TABLE inspection(
    -- 栏目列表
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- '自增ID'
    name varchar(255)  NOT NULL, --  '巡检name: mysql 等',
    count INTEGER  NULL DEFAULT NULL, --  '巡检总数',
    success_count  INTEGER NULL DEFAULT 0, --  '巡检成功数',
    failed_count  INTEGER NULL DEFAULT 0, --  '创建的用户ID',
    connection bool NULL,  --  '连接状态',
    availability  bool NULL,  --  '应用可用性',
    start_time datetime(0) NULL DEFAULT NULL, --  '巡检开始时间',
    end_time datetime(0) NULL DEFAULT NULL --  '巡检结束时间'
);