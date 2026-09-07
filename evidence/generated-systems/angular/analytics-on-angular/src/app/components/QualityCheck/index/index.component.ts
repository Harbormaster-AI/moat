
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { QualityCheckService } from '../../../services/QualityCheck.service';
import { QualityCheck } from '../../../models/QualityCheck';

@Component({
    selector: 'app-index-qualityCheck',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexQualityCheckComponent implements OnInit {

    qualityChecks: QualityCheck[] = [];

    constructor(
        private router: Router,
        private service: QualityCheckService
) {}

    ngOnInit(): void {
        this.getQualityChecks();
}

    getQualityChecks(): void {
        this.service.getQualityChecks().subscribe((res) => {
        this.qualityChecks = res;
    });
}

    deleteQualityCheck(id: any): void {
        this.service.deleteQualityCheck(id)
            .subscribe(() => {
                this.getQualityChecks();
            });
    }
}