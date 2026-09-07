
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateUsageLimitComponent } from './create.component';
import { UsageLimitService } from '../../../services/UsageLimit.service';
import { Router } from '@angular/router';

describe('CreateUsageLimitComponent', () => {
  let component: CreateUsageLimitComponent;
  let fixture: ComponentFixture<CreateUsageLimitComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateUsageLimitComponent
      ],
      providers: [
        UsageLimitService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateUsageLimitComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});