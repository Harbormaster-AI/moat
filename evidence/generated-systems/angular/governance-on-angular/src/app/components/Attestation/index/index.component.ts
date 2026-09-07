
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AttestationService } from '../../../services/Attestation.service';
import { Attestation } from '../../../models/Attestation';

@Component({
    selector: 'app-index-attestation',
    standalone: false,
    templateUrl: './index.component.html',
    styleUrls: ['./index.component.css']
})
export class IndexAttestationComponent implements OnInit {

    attestations: Attestation[] = [];

    constructor(
        private router: Router,
        private service: AttestationService
) {}

    ngOnInit(): void {
        this.getAttestations();
}

    getAttestations(): void {
        this.service.getAttestations().subscribe((res) => {
        this.attestations = res;
    });
}

    deleteAttestation(id: any): void {
        this.service.deleteAttestation(id)
            .subscribe(() => {
                this.getAttestations();
            });
    }
}