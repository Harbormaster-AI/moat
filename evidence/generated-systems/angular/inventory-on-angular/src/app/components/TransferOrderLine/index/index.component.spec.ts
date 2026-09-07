
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTransferOrderLineComponent } from './index.component';
import { TransferOrderLineService } from '../../../services/TransferOrderLine.service';

describe('IndexTransferOrderLineComponent', () => {
  let component: IndexTransferOrderLineComponent;
  let fixture: ComponentFixture<IndexTransferOrderLineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTransferOrderLineComponent
      ],
      providers: [
        TransferOrderLineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTransferOrderLineComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});