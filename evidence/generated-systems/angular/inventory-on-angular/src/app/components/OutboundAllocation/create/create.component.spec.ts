
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateOutboundAllocationComponent } from './create.component';
import { OutboundAllocationService } from '../../../services/OutboundAllocation.service';
import { Router } from '@angular/router';

describe('CreateOutboundAllocationComponent', () => {
  let component: CreateOutboundAllocationComponent;
  let fixture: ComponentFixture<CreateOutboundAllocationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateOutboundAllocationComponent
      ],
      providers: [
        OutboundAllocationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateOutboundAllocationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});