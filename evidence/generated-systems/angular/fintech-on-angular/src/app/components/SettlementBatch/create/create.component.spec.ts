
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSettlementBatchComponent } from './create.component';
import { SettlementBatchService } from '../../../services/SettlementBatch.service';
import { Router } from '@angular/router';

describe('CreateSettlementBatchComponent', () => {
  let component: CreateSettlementBatchComponent;
  let fixture: ComponentFixture<CreateSettlementBatchComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSettlementBatchComponent
      ],
      providers: [
        SettlementBatchService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSettlementBatchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});