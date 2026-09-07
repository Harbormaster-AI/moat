
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ProcedureService } from '../../../services/Procedure.service';
import { Procedure } from '../../../models/Procedure';

@Component({
    selector: 'app-index-procedure',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexProcedureComponent implements OnInit {

    procedures: Procedure[] = [];

    constructor(
        private router: Router,
        private service: ProcedureService
) {}

    ngOnInit(): void {
        this.getProcedures();
}

    getProcedures(): void {
        this.service.getProcedures().subscribe((res) => {
        this.procedures = res;
    });
}

    deleteProcedure(id: any): void {
        this.service.deleteProcedure(id)
            .subscribe(() => {
                this.getProcedures();
            });
    }
}