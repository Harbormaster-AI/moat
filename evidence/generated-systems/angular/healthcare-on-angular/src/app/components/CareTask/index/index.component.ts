
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CareTaskService } from '../../../services/CareTask.service';
import { CareTask } from '../../../models/CareTask';

@Component({
    selector: 'app-index-careTask',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCareTaskComponent implements OnInit {

    careTasks: CareTask[] = [];

    constructor(
        private router: Router,
        private service: CareTaskService
) {}

    ngOnInit(): void {
        this.getCareTasks();
}

    getCareTasks(): void {
        this.service.getCareTasks().subscribe((res) => {
        this.careTasks = res;
    });
}

    deleteCareTask(id: any): void {
        this.service.deleteCareTask(id)
            .subscribe(() => {
                this.getCareTasks();
            });
    }
}