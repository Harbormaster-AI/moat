
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexOutboundAllocationComponent } from './index.component';
import { OutboundAllocationService } from '../../../services/OutboundAllocation.service';

describe('IndexOutboundAllocationComponent', () => {
  let component: IndexOutboundAllocationComponent;
  let fixture: ComponentFixture<IndexOutboundAllocationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexOutboundAllocationComponent
      ],
      providers: [
        OutboundAllocationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexOutboundAllocationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});