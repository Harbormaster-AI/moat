
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInvoiceComponent } from './index.component';
import { InvoiceService } from '../../../services/Invoice.service';

describe('IndexInvoiceComponent', () => {
  let component: IndexInvoiceComponent;
  let fixture: ComponentFixture<IndexInvoiceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInvoiceComponent
      ],
      providers: [
        InvoiceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInvoiceComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});