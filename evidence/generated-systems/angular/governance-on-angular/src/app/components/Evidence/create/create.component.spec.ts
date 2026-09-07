
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateEvidenceComponent } from './create.component';
import { EvidenceService } from '../../../services/Evidence.service';
import { Router } from '@angular/router';

describe('CreateEvidenceComponent', () => {
  let component: CreateEvidenceComponent;
  let fixture: ComponentFixture<CreateEvidenceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateEvidenceComponent
      ],
      providers: [
        EvidenceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateEvidenceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});