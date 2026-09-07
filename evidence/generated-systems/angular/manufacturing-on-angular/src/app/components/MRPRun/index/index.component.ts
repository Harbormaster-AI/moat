
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { MRPRunService } from '../../../services/MRPRun.service';
import { MRPRun } from '../../../models/MRPRun';

@Component({
    selector: 'app-index-mRPRun',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexMRPRunComponent implements OnInit {

    mRPRuns: MRPRun[] = [];

    constructor(
        private router: Router,
        private service: MRPRunService
) {}

    ngOnInit(): void {
        this.getMRPRuns();
}

    getMRPRuns(): void {
        this.service.getMRPRuns().subscribe((res) => {
        this.mRPRuns = res;
    });
}

    deleteMRPRun(id: any): void {
        this.service.deleteMRPRun(id)
            .subscribe(() => {
                this.getMRPRuns();
            });
    }
}