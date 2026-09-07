import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CandidateService } from '../../../services/Candidate.service';
import { Candidate } from '../../../models/Candidate';
import { SubBaseComponent } from '../../Candidate/sub.base.component';

@Component({
    selector: 'app-create-candidate',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCandidateComponent extends SubBaseComponent implements OnInit {

    title = 'Add Candidate';

    candidateForm: FormGroup;
    candidate: Candidate;

    constructor( http: HttpClient,
        private candidateService: CandidateService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.candidateForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      email: ['', Validators.required],
      phone: ['', Validators.required],
      Applications: ['', ],
      Interviews: ['', ],
      Offers: ['', ],
      Documents: ['', ],
      Source: ['', ]
        });
    }

    
    addCandidate(name, email, phone, Applications, Interviews, Offers, Documents, Source): void {
        this.candidateService
        .addCandidate(name, email, phone, Applications, Interviews, Offers, Documents, Source)
            .subscribe(() => {
                this.router.navigate(['/indexCandidate']);
            });
    }

    ngOnInit(): void {
    }
}