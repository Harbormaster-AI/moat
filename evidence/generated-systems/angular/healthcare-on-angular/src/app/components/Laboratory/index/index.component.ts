
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LaboratoryService } from '../../../services/Laboratory.service';
import { Laboratory } from '../../../models/Laboratory';

@Component({
    selector: 'app-index-laboratory',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexLaboratoryComponent implements OnInit {

    laboratorys: Laboratory[] = [];

    constructor(
        private router: Router,
        private service: LaboratoryService
) {}

    ngOnInit(): void {
        this.getLaboratorys();
}

    getLaboratorys(): void {
        this.service.getLaboratorys().subscribe((res) => {
        this.laboratorys = res;
    });
}

    deleteLaboratory(id: any): void {
        this.service.deleteLaboratory(id)
            .subscribe(() => {
                this.getLaboratorys();
            });
    }
}