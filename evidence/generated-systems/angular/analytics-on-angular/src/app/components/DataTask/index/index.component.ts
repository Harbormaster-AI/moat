
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DataTaskService } from '../../../services/DataTask.service';
import { DataTask } from '../../../models/DataTask';

@Component({
    selector: 'app-index-dataTask',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDataTaskComponent implements OnInit {

    dataTasks: DataTask[] = [];

    constructor(
        private router: Router,
        private service: DataTaskService
) {}

    ngOnInit(): void {
        this.getDataTasks();
}

    getDataTasks(): void {
        this.service.getDataTasks().subscribe((res) => {
        this.dataTasks = res;
    });
}

    deleteDataTask(id: any): void {
        this.service.deleteDataTask(id)
            .subscribe(() => {
                this.getDataTasks();
            });
    }
}