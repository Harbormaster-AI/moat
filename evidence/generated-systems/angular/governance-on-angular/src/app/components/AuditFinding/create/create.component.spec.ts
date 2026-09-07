
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAuditFindingComponent } from './create.component';
import { AuditFindingService } from '../../../services/AuditFinding.service';
import { Router } from '@angular/router';

describe('CreateAuditFindingComponent', () => {
  let component: CreateAuditFindingComponent;
  let fixture: ComponentFixture<CreateAuditFindingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAuditFindingComponent
      ],
      providers: [
        AuditFindingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAuditFindingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});