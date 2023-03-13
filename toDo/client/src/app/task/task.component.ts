import {Component, EventEmitter, Input, OnChanges, OnDestroy, OnInit, Output, SimpleChanges} from '@angular/core';
import {Task} from "../models";

@Component({
  selector: 'app-task',
  templateUrl: './task.component.html',
  styleUrls: ['./task.component.css']
})
export class TaskComponent{
  @Input() task: Task;
  @Output() DeleteTask = new EventEmitter();
  @Output() SaveTask = new EventEmitter();
  edit: boolean;
  message: string;

  constructor() {
    this.task = {} as Task;
    this.edit = this.task.complete;
    this.message = this.task.message;
  }








  deleteTask(id:number){
    this.DeleteTask.emit(id);
  }

  saveTask(task: Task){
    this.edit = !this.edit;
    if(this.message){
      task.message = this.message;
    }
    this.SaveTask.emit(task);
    this.message = this.task.message;
  }
  changeTask(task: Task){
    this.SaveTask.emit(task);
    this.message = this.task.message;
  }
  formInit(){
    this.message = this.task.message;
    this.edit = true;
  }

}
