import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AdSlotService } from '../../../services/AdSlot.service';
import { AdSlot } from '../../../models/AdSlot';
import { SubBaseComponent } from '../../AdSlot/sub.base.component';

@Component({
    selector: 'app-create-adSlot',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAdSlotComponent extends SubBaseComponent implements OnInit {

    title = 'Add AdSlot';

    adSlotForm: FormGroup;
    adSlot: AdSlot;

    constructor( http: HttpClient,
        private adSlotService: AdSlotService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.adSlotForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  slotCode: ['', Validators.required],
      width: ['', Validators.required],
      height: ['', Validators.required],
      floorPrice: ['', Validators.required],
      InventorySource: ['', ],
      Placements: ['', ],
      Rates: ['', ],
      Format: ['', ]
        });
    }

    
    addAdSlot(slotCode, width, height, floorPrice, InventorySource, Placements, Rates, Format): void {
        this.adSlotService
        .addAdSlot(slotCode, width, height, floorPrice, InventorySource, Placements, Rates, Format)
            .subscribe(() => {
                this.router.navigate(['/indexAdSlot']);
            });
    }

    ngOnInit(): void {
    }
}