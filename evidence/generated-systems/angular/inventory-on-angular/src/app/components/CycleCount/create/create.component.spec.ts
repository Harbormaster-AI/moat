
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCycleCountComponent } from './create.component';
import { CycleCountService } from '../../../services/CycleCount.service';
import { Router } from '@angular/router';

describe('CreateCycleCountComponent', () => {
  let component: CreateCycleCountComponent;
  let fixture: ComponentFixture<CreateCycleCountComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCycleCountComponent
      ],
      providers: [
        CycleCountService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCycleCountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});