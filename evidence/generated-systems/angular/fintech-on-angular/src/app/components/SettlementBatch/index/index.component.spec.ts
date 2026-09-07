
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSettlementBatchComponent } from './index.component';
import { SettlementBatchService } from '../../../services/SettlementBatch.service';

describe('IndexSettlementBatchComponent', () => {
  let component: IndexSettlementBatchComponent;
  let fixture: ComponentFixture<IndexSettlementBatchComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSettlementBatchComponent
      ],
      providers: [
        SettlementBatchService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSettlementBatchComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});