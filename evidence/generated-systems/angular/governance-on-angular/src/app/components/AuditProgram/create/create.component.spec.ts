
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateAuditProgramComponent } from './create.component';
import { AuditProgramService } from '../../../services/AuditProgram.service';
import { Router } from '@angular/router';

describe('CreateAuditProgramComponent', () => {
  let component: CreateAuditProgramComponent;
  let fixture: ComponentFixture<CreateAuditProgramComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateAuditProgramComponent
      ],
      providers: [
        AuditProgramService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateAuditProgramComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});