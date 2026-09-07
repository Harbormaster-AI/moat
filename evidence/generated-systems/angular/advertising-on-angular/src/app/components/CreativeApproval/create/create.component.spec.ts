
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCreativeApprovalComponent } from './create.component';
import { CreativeApprovalService } from '../../../services/CreativeApproval.service';
import { Router } from '@angular/router';

describe('CreateCreativeApprovalComponent', () => {
  let component: CreateCreativeApprovalComponent;
  let fixture: ComponentFixture<CreateCreativeApprovalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCreativeApprovalComponent
      ],
      providers: [
        CreativeApprovalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCreativeApprovalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});