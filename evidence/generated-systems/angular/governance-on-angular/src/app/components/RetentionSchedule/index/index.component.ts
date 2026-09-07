
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RetentionScheduleService } from '../../../services/RetentionSchedule.service';
import { RetentionSchedule } from '../../../models/RetentionSchedule';

@Component({
    selector: 'app-index-retentionSchedule',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexRetentionScheduleComponent implements OnInit {

    retentionSchedules: RetentionSchedule[] = [];

    constructor(
        private router: Router,
        private service: RetentionScheduleService
) {}

    ngOnInit(): void {
        this.getRetentionSchedules();
}

    getRetentionSchedules(): void {
        this.service.getRetentionSchedules().subscribe((res) => {
        this.retentionSchedules = res;
    });
}

    deleteRetentionSchedule(id: any): void {
        this.service.deleteRetentionSchedule(id)
            .subscribe(() => {
                this.getRetentionSchedules();
            });
    }
}