
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { BuildScheduleService } from '../../../services/BuildSchedule.service';
import { BuildSchedule } from '../../../models/BuildSchedule';

@Component({
    selector: 'app-index-buildSchedule',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexBuildScheduleComponent implements OnInit {

    buildSchedules: BuildSchedule[] = [];

    constructor(
        private router: Router,
        private service: BuildScheduleService
) {}

    ngOnInit(): void {
        this.getBuildSchedules();
}

    getBuildSchedules(): void {
        this.service.getBuildSchedules().subscribe((res) => {
        this.buildSchedules = res;
    });
}

    deleteBuildSchedule(id: any): void {
        this.service.deleteBuildSchedule(id)
            .subscribe(() => {
                this.getBuildSchedules();
            });
    }
}