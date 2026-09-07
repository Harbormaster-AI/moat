import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuditFindingService } from '../../../services/AuditFinding.service';
import { SubBaseComponent } from '../../AuditFinding/sub.base.component';


@Component({
    selector: 'app-edit-auditFinding',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAuditFindingComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AuditFinding';

    auditFindingForm: FormGroup;
    auditFinding: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AuditFindingService,
        private fb: FormBuilder
) {
        super(http);
        this.auditFindingForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      description: ['', Validators.required],
      dueDate: ['', Validators.required],
      Engagement: ['', ],
      Workpaper: ['', ],
      CorrectiveActions: ['', ],
      RelatedRisks: ['', ],
      RelatedControls: ['', ],
      Issues: ['', ],
      Severity: ['', ],
      Status: ['', ]
        });
    }

    
    updateAuditFinding(title, description, dueDate, Engagement, Workpaper, CorrectiveActions, RelatedRisks, RelatedControls, Issues, Severity, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAuditFinding(title, description, dueDate, Engagement, Workpaper, CorrectiveActions, RelatedRisks, RelatedControls, Issues, Severity, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAuditFinding']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAuditFinding(params['id']).subscribe(res => {
                this.auditFinding = res;
            });
        });
    }
}