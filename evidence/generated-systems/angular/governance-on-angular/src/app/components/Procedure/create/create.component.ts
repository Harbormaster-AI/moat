import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProcedureService } from '../../../services/Procedure.service';
import { Procedure } from '../../../models/Procedure';
import { SubBaseComponent } from '../../Procedure/sub.base.component';

@Component({
    selector: 'app-create-procedure',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateProcedureComponent extends SubBaseComponent implements OnInit {

    title = 'Add Procedure';

    procedureForm: FormGroup;
    procedure: Procedure;

    constructor( http: HttpClient,
        private procedureService: ProcedureService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.procedureForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      versionLabel: ['', Validators.required],
      Policy: ['', ],
      Controls: ['', ],
      Status: ['', ]
        });
    }

    
    addProcedure(title, versionLabel, Policy, Controls, Status): void {
        this.procedureService
        .addProcedure(title, versionLabel, Policy, Controls, Status)
            .subscribe(() => {
                this.router.navigate(['/indexProcedure']);
            });
    }

    ngOnInit(): void {
    }
}