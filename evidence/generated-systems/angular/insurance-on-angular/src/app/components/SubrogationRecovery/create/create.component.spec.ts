
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSubrogationRecoveryComponent } from './create.component';
import { SubrogationRecoveryService } from '../../../services/SubrogationRecovery.service';
import { Router } from '@angular/router';

describe('CreateSubrogationRecoveryComponent', () => {
  let component: CreateSubrogationRecoveryComponent;
  let fixture: ComponentFixture<CreateSubrogationRecoveryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSubrogationRecoveryComponent
      ],
      providers: [
        SubrogationRecoveryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSubrogationRecoveryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});