import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuditProgramService } from '../../../services/AuditProgram.service';
import { SubBaseComponent } from '../../AuditProgram/sub.base.component';


@Component({
    selector: 'app-edit-auditProgram',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAuditProgramComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AuditProgram';

    auditProgramForm: FormGroup;
    auditProgram: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AuditProgramService,
        private fb: FormBuilder
) {
        super(http);
        this.auditProgramForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      scope: ['', Validators.required],
      Organization: ['', ],
      Engagements: ['', ],
      Cycle: ['', ],
      Status: ['', ]
        });
    }

    
    updateAuditProgram(name, scope, Organization, Engagements, Cycle, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAuditProgram(name, scope, Organization, Engagements, Cycle, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAuditProgram']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAuditProgram(params['id']).subscribe(res => {
                this.auditProgram = res;
            });
        });
    }
}