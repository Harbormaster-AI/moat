
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateQuarantineComponent } from './create.component';
import { QuarantineService } from '../../../services/Quarantine.service';
import { Router } from '@angular/router';

describe('CreateQuarantineComponent', () => {
  let component: CreateQuarantineComponent;
  let fixture: ComponentFixture<CreateQuarantineComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateQuarantineComponent
      ],
      providers: [
        QuarantineService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateQuarantineComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});