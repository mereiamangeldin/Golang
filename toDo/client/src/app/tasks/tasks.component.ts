import {Component, OnInit} from '@angular/core';
import {Task} from "../models";
import {HttpClient} from "@angular/common/http";

@Component({
  selector: 'app-tasks',
  templateUrl: './tasks.component.html',
  styleUrls: ['./tasks.component.css']
})
export class TasksComponent implements OnInit{
  tasks: Task[] | undefined = [];

  message: string;
  constructor(
    private httpClient: HttpClient
  ) {
    this.message = "";

  }

  async AddTask(){
    if (this.message){
      const newTask = {id: 0, message: this.message, complete: false};
      await this.httpClient.post('api', newTask).toPromise();
      await this.loadTasks();
      this.message = "";
    }
  }
  async DeleteTask(id:number){
    await this.httpClient.delete(`api/${id}`).toPromise();
    await this.loadTasks();
  }
  async SaveTask(task: Task){
    await this.httpClient.put(`api/${task.ID}`, task).toPromise();
    await this.loadTasks();
  }


  async loadTasks(){
    this.tasks = await this.httpClient.get<Task[]>('api').toPromise()
  }

  async ngOnInit() {
    await this.loadTasks()
  }
}


