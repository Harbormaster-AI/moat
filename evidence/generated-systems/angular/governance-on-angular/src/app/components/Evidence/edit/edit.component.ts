import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { EvidenceService } from '../../../services/Evidence.service';
import { SubBaseComponent } from '../../Evidence/sub.base.component';


@Component({
    selector: 'app-edit-evidence',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditEvidenceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Evidence';

    evidenceForm: FormGroup;
    evidence: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: EvidenceService,
        private fb: FormBuilder
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

    
    updateEvidence(title, locationUrl, receivedDate, ControlTest, Control, Obligation, Workpaper, EvidenceType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateEvidence(title, locationUrl, receivedDate, ControlTest, Control, Obligation, Workpaper, EvidenceType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexEvidence']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getEvidence(params['id']).subscribe(res => {
                this.evidence = res;
            });
        });
    }
}