import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuditFindingService } from '../../../services/AuditFinding.service';
import { AuditFinding } from '../../../models/AuditFinding';
import { SubBaseComponent } from '../../AuditFinding/sub.base.component';

@Component({
    selector: 'app-create-auditFinding',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAuditFindingComponent extends SubBaseComponent implements OnInit {

    title = 'Add AuditFinding';

    auditFindingForm: FormGroup;
    auditFinding: AuditFinding;

    constructor( http: HttpClient,
        private auditFindingService: AuditFindingService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAuditFinding(title, description, dueDate, Engagement, Workpaper, CorrectiveActions, RelatedRisks, RelatedControls, Issues, Severity, Status): void {
        this.auditFindingService
        .addAuditFinding(title, description, dueDate, Engagement, Workpaper, CorrectiveActions, RelatedRisks, RelatedControls, Issues, Severity, Status)
            .subscribe(() => {
                this.router.navigate(['/indexAuditFinding']);
            });
    }

    ngOnInit(): void {
    }
}