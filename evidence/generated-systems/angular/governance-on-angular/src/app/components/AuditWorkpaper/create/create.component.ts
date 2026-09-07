import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AuditWorkpaperService } from '../../../services/AuditWorkpaper.service';
import { AuditWorkpaper } from '../../../models/AuditWorkpaper';
import { SubBaseComponent } from '../../AuditWorkpaper/sub.base.component';

@Component({
    selector: 'app-create-auditWorkpaper',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAuditWorkpaperComponent extends SubBaseComponent implements OnInit {

    title = 'Add AuditWorkpaper';

    auditWorkpaperForm: FormGroup;
    auditWorkpaper: AuditWorkpaper;

    constructor( http: HttpClient,
        private auditWorkpaperService: AuditWorkpaperService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.auditWorkpaperForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  workpaperRef: ['', Validators.required],
      subject: ['', Validators.required],
      workpaperUrl: ['', Validators.required],
      Engagement: ['', ],
      Evidence: ['', ],
      Findings: ['', ]
        });
    }

    
    addAuditWorkpaper(workpaperRef, subject, workpaperUrl, Engagement, Evidence, Findings): void {
        this.auditWorkpaperService
        .addAuditWorkpaper(workpaperRef, subject, workpaperUrl, Engagement, Evidence, Findings)
            .subscribe(() => {
                this.router.navigate(['/indexAuditWorkpaper']);
            });
    }

    ngOnInit(): void {
    }
}