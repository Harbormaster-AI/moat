
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { JobRequisitionService } from '../../../services/JobRequisition.service';
import { JobRequisition } from '../../../models/JobRequisition';

@Component({
    selector: 'app-index-jobRequisition',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexJobRequisitionComponent implements OnInit {

    jobRequisitions: JobRequisition[] = [];

    constructor(
        private router: Router,
        private service: JobRequisitionService
) {}

    ngOnInit(): void {
        this.getJobRequisitions();
}

    getJobRequisitions(): void {
        this.service.getJobRequisitions().subscribe((res) => {
        this.jobRequisitions = res;
    });
}

    deleteJobRequisition(id: any): void {
        this.service.deleteJobRequisition(id)
            .subscribe(() => {
                this.getJobRequisitions();
            });
    }
}