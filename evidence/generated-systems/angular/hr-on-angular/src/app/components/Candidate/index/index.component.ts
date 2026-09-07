
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { CandidateService } from '../../../services/Candidate.service';
import { Candidate } from '../../../models/Candidate';

@Component({
    selector: 'app-index-candidate',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexCandidateComponent implements OnInit {

    candidates: Candidate[] = [];

    constructor(
        private router: Router,
        private service: CandidateService
) {}

    ngOnInit(): void {
        this.getCandidates();
}

    getCandidates(): void {
        this.service.getCandidates().subscribe((res) => {
        this.candidates = res;
    });
}

    deleteCandidate(id: any): void {
        this.service.deleteCandidate(id)
            .subscribe(() => {
                this.getCandidates();
            });
    }
}