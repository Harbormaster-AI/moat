import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EvidenceService } from '../../../services/Evidence.service';
import { Evidence } from '../../../models/Evidence';
import { SubBaseComponent } from '../../Evidence/sub.base.component';

@Component({
    selector: 'app-create-evidence',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateEvidenceComponent extends SubBaseComponent implements OnInit {

    title = 'Add Evidence';

    evidenceForm: FormGroup;
    evidence: Evidence;

    constructor( http: HttpClient,
        private evidenceService: EvidenceService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.evidenceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      locationUrl: ['', Validators.required],
      receivedDate: ['', Validators.required],
      ControlTest: ['', ],
      Control: ['', ],
      Obligation: ['', ],
      Workpaper: ['', ],
      EvidenceType: ['', ]
        });
    }

    
    addEvidence(title, locationUrl, receivedDate, ControlTest, Control, Obligation, Workpaper, EvidenceType): void {
        this.evidenceService
        .addEvidence(title, locationUrl, receivedDate, ControlTest, Control, Obligation, Workpaper, EvidenceType)
            .subscribe(() => {
                this.router.navigate(['/indexEvidence']);
            });
    }

    ngOnInit(): void {
    }
}