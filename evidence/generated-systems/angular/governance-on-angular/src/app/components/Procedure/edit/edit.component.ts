import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ProcedureService } from '../../../services/Procedure.service';
import { SubBaseComponent } from '../../Procedure/sub.base.component';


@Component({
    selector: 'app-edit-procedure',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditProcedureComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Procedure';

    procedureForm: FormGroup;
    procedure: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ProcedureService,
        private fb: FormBuilder
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

    
    updateProcedure(title, versionLabel, Policy, Controls, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateProcedure(title, versionLabel, Policy, Controls, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexProcedure']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getProcedure(params['id']).subscribe(res => {
                this.procedure = res;
            });
        });
    }
}