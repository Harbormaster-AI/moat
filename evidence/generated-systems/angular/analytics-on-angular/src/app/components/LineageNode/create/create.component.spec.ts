
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLineageNodeComponent } from './create.component';
import { LineageNodeService } from '../../../services/LineageNode.service';
import { Router } from '@angular/router';

describe('CreateLineageNodeComponent', () => {
  let component: CreateLineageNodeComponent;
  let fixture: ComponentFixture<CreateLineageNodeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLineageNodeComponent
      ],
      providers: [
        LineageNodeService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLineageNodeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});