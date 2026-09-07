
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { DiagnosisService } from '../../../services/Diagnosis.service';
import { Diagnosis } from '../../../models/Diagnosis';

@Component({
    selector: 'app-index-diagnosis',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexDiagnosisComponent implements OnInit {

    diagnosiss: Diagnosis[] = [];

    constructor(
        private router: Router,
        private service: DiagnosisService
) {}

    ngOnInit(): void {
        this.getDiagnosiss();
}

    getDiagnosiss(): void {
        this.service.getDiagnosiss().subscribe((res) => {
        this.diagnosiss = res;
    });
}

    deleteDiagnosis(id: any): void {
        this.service.deleteDiagnosis(id)
            .subscribe(() => {
                this.getDiagnosiss();
            });
    }
}