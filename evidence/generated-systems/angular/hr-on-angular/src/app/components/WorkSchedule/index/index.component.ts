
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { WorkScheduleService } from '../../../services/WorkSchedule.service';
import { WorkSchedule } from '../../../models/WorkSchedule';

@Component({
    selector: 'app-index-workSchedule',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexWorkScheduleComponent implements OnInit {

    workSchedules: WorkSchedule[] = [];

    constructor(
        private router: Router,
        private service: WorkScheduleService
) {}

    ngOnInit(): void {
        this.getWorkSchedules();
}

    getWorkSchedules(): void {
        this.service.getWorkSchedules().subscribe((res) => {
        this.workSchedules = res;
    });
}

    deleteWorkSchedule(id: any): void {
        this.service.deleteWorkSchedule(id)
            .subscribe(() => {
                this.getWorkSchedules();
            });
    }
}