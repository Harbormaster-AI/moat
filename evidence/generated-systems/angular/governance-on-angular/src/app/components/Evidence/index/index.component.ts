
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { EvidenceService } from '../../../services/Evidence.service';
import { Evidence } from '../../../models/Evidence';

@Component({
    selector: 'app-index-evidence',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexEvidenceComponent implements OnInit {

    evidences: Evidence[] = [];

    constructor(
        private router: Router,
        private service: EvidenceService
) {}

    ngOnInit(): void {
        this.getEvidences();
}

    getEvidences(): void {
        this.service.getEvidences().subscribe((res) => {
        this.evidences = res;
    });
}

    deleteEvidence(id: any): void {
        this.service.deleteEvidence(id)
            .subscribe(() => {
                this.getEvidences();
            });
    }
}