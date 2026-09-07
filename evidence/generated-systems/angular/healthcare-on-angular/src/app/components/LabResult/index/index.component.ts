
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LabResultService } from '../../../services/LabResult.service';
import { LabResult } from '../../../models/LabResult';

@Component({
    selector: 'app-index-labResult',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLabResultComponent implements OnInit {

    labResults: LabResult[] = [];

    constructor(
        private router: Router,
        private service: LabResultService
) {}

    ngOnInit(): void {
        this.getLabResults();
}

    getLabResults(): void {
        this.service.getLabResults().subscribe((res) => {
        this.labResults = res;
    });
}

    deleteLabResult(id: any): void {
        this.service.deleteLabResult(id)
            .subscribe(() => {
                this.getLabResults();
            });
    }
}