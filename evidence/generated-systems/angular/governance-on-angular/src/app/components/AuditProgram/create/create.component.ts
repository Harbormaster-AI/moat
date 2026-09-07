import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuditProgramService } from '../../../services/AuditProgram.service';
import { AuditProgram } from '../../../models/AuditProgram';
import { SubBaseComponent } from '../../AuditProgram/sub.base.component';

@Component({
    selector: 'app-create-auditProgram',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAuditProgramComponent extends SubBaseComponent implements OnInit {

    title = 'Add AuditProgram';

    auditProgramForm: FormGroup;
    auditProgram: AuditProgram;

    constructor( http: HttpClient,
        private auditProgramService: AuditProgramService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAuditProgram(name, scope, Organization, Engagements, Cycle, Status): void {
        this.auditProgramService
        .addAuditProgram(name, scope, Organization, Engagements, Cycle, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAuditProgram']);
            });
    }

    ngOnInit(): void {
    }
}