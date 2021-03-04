package main

// This code was generated by a tool. Don't modify it manually.
// http://sqldalmaker.sourceforge.net

type TasksDao struct {
    ds *DataStore
}

// (C)RUD: tasks
// Generated values are passed to DTO.
// Returns the number of affected rows or -1 on error.

func (dao *TasksDao) createTask(p *Task) {
    sql := "insert into tasks (g_id, t_priority, t_date, t_subject, t_comments) values (?, ?, ?, ?, ?)"
    res := dao.ds.insert(sql, p.GId, p.TPriority, p.TDate, p.TSubject, p.TComments)
    p.TId = res
}

// C(R)UD: tasks

func (dao *TasksDao) readTask(TId int64) Task {
    sql := "select * from tasks where t_id=?"
    rd := dao.ds.queryRow(sql, TId)
    obj := Task{}
    obj.TId = rd["t_id"].(int64)
    obj.GId = rd["g_id"].(int64)
    obj.TPriority = rd["t_priority"].(int64)
    obj.TDate = rd["t_date"].(string)
    obj.TSubject = rd["t_subject"].(string)
    obj.TComments = rd["t_comments"].(string)
    return obj
}

// CR(U)D: tasks
// Returns the number of affected rows or -1 on error.

func (dao *TasksDao) updateTask(p *Task) int64 {
    sql := "update tasks set g_id=?, t_priority=?, t_date=?, t_subject=?, t_comments=? where t_id=?"
    return dao.ds.execDML(sql, p.GId, p.TPriority, p.TDate, p.TSubject, p.TComments, p.TId)
}

// CRU(D): tasks
// Returns the number of affected rows or -1 on error.

func (dao *TasksDao) deleteTask(TId int64) int64 {
    sql := "delete from tasks where t_id=?"
    return dao.ds.execDML(sql, TId);
}

func (dao *TasksDao) getGroupTasks(gId int64) []Task {
    sql := "select * from tasks where g_id =?" + 
        "\n order by t_id"
    var res []Task
    onDto := func(rd map[string]interface{}) {
        obj := Task{}
        obj.TId = rd["t_id"].(int64)
        obj.GId = rd["g_id"].(int64)
        obj.TPriority = rd["t_priority"].(int64)
        obj.TDate = rd["t_date"].(string)
        obj.TSubject = rd["t_subject"].(string)
        obj.TComments = rd["t_comments"].(string)
        res = append(res, obj)
    }
    dao.ds.queryAllRows(sql, onDto, gId)
    return res
}

// Returns the number of affected rows or -1 on error.

func (dao *TasksDao) deleteGroupTasks(gId string) int64 {
    sql := "delete from tasks where g_id=?"
    return dao.ds.execDML(sql, gId);
}

func (dao *TasksDao) getCount() interface{} {
    sql := "select count(*) from tasks"
    return dao.ds.query(sql).(interface{})
}
