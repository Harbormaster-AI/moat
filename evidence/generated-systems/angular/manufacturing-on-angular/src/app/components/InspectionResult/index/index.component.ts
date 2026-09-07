
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { InspectionResultService } from '../../../services/InspectionResult.service';
import { InspectionResult } from '../../../models/InspectionResult';

@Component({
    selector: 'app-index-inspectionResult',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexInspectionResultComponent implements OnInit {

    inspectionResults: InspectionResult[] = [];

    constructor(
        private router: Router,
        private service: InspectionResultService
) {}

    ngOnInit(): void {
        this.getInspectionResults();
}

    getInspectionResults(): void {
        this.service.getInspectionResults().subscribe((res) => {
        this.inspectionResults = res;
    });
}

    deleteInspectionResult(id: any): void {
        this.service.deleteInspectionResult(id)
            .subscribe(() => {
                this.getInspectionResults();
            });
    }
}