
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPurchaseOrderLineComponent } from './index.component';
import { PurchaseOrderLineService } from '../../../services/PurchaseOrderLine.service';

describe('IndexPurchaseOrderLineComponent', () => {
  let component: IndexPurchaseOrderLineComponent;
  let fixture: ComponentFixture<IndexPurchaseOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPurchaseOrderLineComponent
      ],
      providers: [
        PurchaseOrderLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPurchaseOrderLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});